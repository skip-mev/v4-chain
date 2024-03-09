package client

import (
	"context"
	"fmt"
	"sync"

	"cosmossdk.io/log"
	"google.golang.org/grpc"

	oracletypes "github.com/skip-mev/slinky/pkg/types"

	appflags "github.com/dydxprotocol/v4-chain/protocol/app/flags"
	daemontypes "github.com/dydxprotocol/v4-chain/protocol/daemons/types"
	"github.com/dydxprotocol/v4-chain/protocol/lib/slinky"
	pricetypes "github.com/dydxprotocol/v4-chain/protocol/x/prices/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// MarketPairFetcher is a lightweight process run in a goroutine by the slinky client.
// Its purpose is to periodically query the prices module for MarketParams and create
// an easily indexed mapping between Slinky's CurrencyPair type and the corresponding ID
// in the x/prices module.
type MarketPairFetcher interface {
	Start(context.Context, appflags.Flags, daemontypes.GrpcClient) error
	Stop()
	GetIDForPair(oracletypes.CurrencyPair) (uint32, error)
	FetchIdMappings(context.Context) error
}

// MarketPairFetcherImpl implements the MarketPairFetcher interface.
type MarketPairFetcherImpl struct {
	Logger            log.Logger
	QueryConn         *grpc.ClientConn
	PricesQueryClient pricetypes.QueryClient

	// compatMappings stores a mapping between CurrencyPair and the corresponding market(param|price) ID
	compatMappings map[oracletypes.CurrencyPair]uint32
	compatMu       sync.RWMutex
}

func NewMarketPairFetcher(logger log.Logger) MarketPairFetcher {
	return &MarketPairFetcherImpl{
		Logger:         logger,
		compatMappings: make(map[oracletypes.CurrencyPair]uint32),
	}
}

// Start opens the grpc connections for the fetcher.
func (m *MarketPairFetcherImpl) Start(
	ctx context.Context,
	appFlags appflags.Flags,
	grpcClient daemontypes.GrpcClient) error {
	// Create the query client connection
	queryConn, err := grpcClient.NewTcpConnection(ctx, appFlags.GrpcAddress)
	if err != nil {
		m.Logger.Error(
			"Failed to establish gRPC connection",
			"gRPC address", appFlags.GrpcAddress,
			"error", err,
		)
		return err
	}
	m.PricesQueryClient = pricetypes.NewQueryClient(queryConn)
	return nil
}

// Stop closes all existing connections.
func (m *MarketPairFetcherImpl) Stop() {
	if m.QueryConn != nil {
		_ = m.QueryConn.Close()
	}
}

// GetIDForPair returns the ID corresponding to the currency pair in the x/prices module.
// If the currency pair is not found it will return an error.
func (m *MarketPairFetcherImpl) GetIDForPair(cp oracletypes.CurrencyPair) (uint32, error) {
	var id uint32
	m.compatMu.RLock()
	defer m.compatMu.RUnlock()
	id, found := m.compatMappings[cp]
	if !found {
		return id, fmt.Errorf("pair %s not found in compatMappings", cp.String())
	}
	return id, nil
}

// FetchIdMappings is run periodically to refresh the cache of known mappings between
// CurrencyPair and MarketParam ID.
func (m *MarketPairFetcherImpl) FetchIdMappings(ctx context.Context) error {
	params := make([]pricetypes.MarketParam, 0)

	// Fetch all market params
	if err := MakePaginatedQuery(ctx, func(ctx context.Context, req *query.PageRequest) (PaginatedGRPCResponse, error) {
		res, err := m.PricesQueryClient.AllMarketParams(ctx, &pricetypes.QueryAllMarketParamsRequest{
			Pagination: req,
		})
		if err != nil {
			return nil, err
		}
		params = append(params, res.MarketParams...)

		return res, nil
	}); err != nil {
		return err
	}

	var compatMappings = make(map[oracletypes.CurrencyPair]uint32, len(params))
	for _, mp := range params {
		cp, err := slinky.MarketPairToCurrencyPair(mp.Pair)
		if err != nil {
			return err
		}
		m.Logger.Debug("Mapped market to pair", "market id", mp.Id, "currency pair", cp.String())
		compatMappings[cp] = mp.Id
	}
	m.compatMu.Lock()
	defer m.compatMu.Unlock()
	m.compatMappings = compatMappings
	return nil
}

type PaginatedGRPCResponse interface {
	GetPagination() (*query.PageResponse)
}

type QueryFunction func(ctx context.Context, req *query.PageRequest) (PaginatedGRPCResponse, error)

// PaginatedRequestHandler
func MakePaginatedQuery(ctx context.Context, qfn QueryFunction) error {
	pag := &query.PageRequest{
		Limit: 10000, 
	}

	for {
		// Get the next page of results
		resp, err := qfn(ctx, pag)
		if err != nil {
			return err
		}

		// Process the results
		if len(resp.GetPagination().NextKey) == 0 {
			break
		}

		// Update the page request
		pag.Key = resp.GetPagination().NextKey
	}
	return nil
}