package petri

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/dydxprotocol/v4-chain/protocol/testutil/encoding"
	pricestypes "github.com/dydxprotocol/v4-chain/protocol/x/prices/types"
	"github.com/skip-mev/petri/core/v2/provider"
	"github.com/skip-mev/petri/core/v2/provider/docker"
	"github.com/skip-mev/petri/core/v2/provider/digitalocean"
	petritypes "github.com/skip-mev/petri/core/v2/types"
	"github.com/skip-mev/petri/cosmos/v2/chain"
	"github.com/skip-mev/petri/cosmos/v2/node"
	"go.uber.org/zap"
	"strconv"
)

const (
	denom            = "dv4tnt"
	prefix           = "dydx"
	homeDir          = "/petri-test"
	appConfigPath    = "config/app.toml"
	cometConfigPath  = "config/config.toml"
	oracleConfigPath = "oracle.json"
	marketConfigPath = "markets.json"
	oraclePort       = "8080"
	oracleMetricsPort = "8010"
	appOracleMetricsPort = "26661"
	cometMetricsPort = "26660"
	cometProfilerPort = "6060"
	envProviderType = "PETRI_LOAD_TEST_PROVIDER_TYPE"
	envDigitalOceanAPIKey = "PETRI_LOAD_TEST_DIGITAL_OCEAN_API_KEY"
	digitalOceanProviderType = "digitalocean"
	envDigitalOceanImageID = "PETRI_LOAD_TEST_DIGITAL_OCEAN_IMAGE_ID"
	envNumMarkets = "PETRI_LOAD_TEST_NUM_MARKETS"
	dockerProviderType = "docker"
	priceDaemonEnabled = "price-daemon-enabled"
	bridgeDaemonEnabled = "bridge-daemon-enabled"
	liquidationDaemonEnabled = "liquidation-daemon-enabled"
	slinkyDaemonEnabled = "slinky-daemon-enabled"
	url = "https://api.gateio.ws/api/v4/spot/currency_pairs"
)
var doRegions = []string{"blr1", "blr1", "lon1", "ams3"}

func GetChainConfig() (petritypes.ChainConfig, error) {
	// get the digital ocean image ID
	doEnabled := os.Getenv(envProviderType) == digitalOceanProviderType
	var doImageID int
	if doEnabled {
		var err error
		doImageID, err = strconv.Atoi(os.Getenv(envDigitalOceanImageID))
		if err != nil {
			return petritypes.ChainConfig{}, fmt.Errorf("failed to parse digital ocean image ID: %w", err)
		}
	}

	return petritypes.ChainConfig{
		Denom:         denom,
		Decimals:      6,
		NumValidators: 4,
		NumNodes:      2,
		BinaryName:    "dydxprotocold",
		Image: provider.ImageDefinition{
			Image: "docker.io/nikhilv01/dydxprotocol-base:latest",
			UID:   "1000",
			GID:   "1000",
		},
		SidecarImage: provider.ImageDefinition{
			Image: "docker.io/nikhilv01/dydxprotocol-base:latest",
			UID:   "1000",
			GID:   "1000",
		},
		SidecarArgs:    []string{
			"slinky", 
			"--oracle-config-path", fmt.Sprintf("/etc/%s", oracleConfigPath), 
			"--market-config-path", fmt.Sprintf("/etc/%s", marketConfigPath), 
			"-host", "0.0.0.0", 
			"-port", "8080",
		},
		GasPrices:      "0dv4tnt",
		GasAdjustment:  1.5,
		Bech32Prefix:   prefix,
		EncodingConfig: encoding.GetTestEncodingCfg(),
		HomeDir:        homeDir,
		SidecarHomeDir: "/etc",
		SidecarPorts:   []string{oraclePort, oracleMetricsPort},
		CoinType:       "118",
		ChainId:        "dydx-1",
		ModifyGenesis:  GetGenesisModifier(),
		WalletConfig: petritypes.WalletConfig{
			DerivationFn:     hd.Secp256k1.Derive(),
			GenerationFn:     hd.Secp256k1.Generate(),
			Bech32Prefix:     prefix,
			HDPath:           hd.CreateHDPath(0, 0, 0),
			SigningAlgorithm: "secp256k1",
		},
		NodeCreator: node.CreateNode, // modify to account for additional parameters
		GenesisDelegation: big.NewInt(10_000_000_000_000),
		GenesisBalance:    big.NewInt(100_000_000_000_000),
		NodeDefinitionModifier: func(def provider.TaskDefinition, nodeConfig petritypes.NodeConfig) provider.TaskDefinition {
			// update flags
			def.Entrypoint = append(def.Entrypoint, []string{
				"--price-daemon-enabled=false",
				"--bridge-daemon-enabled=false",
				"--liquidation-daemon-enabled=false",
				"--slinky-daemon-enabled=true",
				}...
			)

			if doEnabled {
				def.ProviderSpecificConfig = digitalocean.DigitalOceanTaskConfig{
					Size: "c-16",
					Region: doRegions[nodeConfig.Index%len(doRegions)], // multiplex onto multiple regions
					ImageID: doImageID,
				}

				// update the sidecar Provider specific configs as well
				for i := range def.Sidecars {
					def.Sidecars[i].ProviderSpecificConfig = def.ProviderSpecificConfig
				}
			}
			return def
		},
	}, nil
}

func GetGenesisModifier() petritypes.GenesisModifier {
	var genKVs = []chain.GenesisKV{
		{
			Key:   "app_state.gov.params.voting_period",
			Value: "10s",
		},
		{
			Key:   "app_state.gov.params.expedited_voting_period",
			Value: "5s",
		},
		{
			Key:   "app_state.gov.params.max_deposit_period",
			Value: "1s",
		},
		{
			Key:   "app_state.gov.params.min_deposit.0.denom",
			Value: denom,
		},
		{
			Key:   "app_state.gov.params.min_deposit.0.amount",
			Value: "1",
		},
		{
			Key:   "app_state.gov.params.threshold",
			Value: "0.1",
		},
		{
			Key:   "app_state.gov.params.quorum",
			Value: "0",
		},
		{
			Key:   "consensus.params.abci.vote_extensions_enable_height",
			Value: "2",
		},
		{
			Key:   "app_state.staking.params.bond_denom",
			Value: denom,
		},
	}

	// update all currency-pairs
	cps, err := getCPsFromGate(url)
	if err != nil {
		panic(err)
	}

	marketParams := make([]pricestypes.MarketParam, len(cps))
	marketPrices := make([]pricestypes.MarketPrice, len(cps))
	for i, cp := range cps {
		marketParams[i] = pricestypes.MarketParam{
			Id: uint32(i),
			Pair: fmt.Sprintf("%s-%s", cp.Base, cp.Quote),
			Exponent: -8,
			MinExchanges: 1,
			MinPriceChangePpm: 1,
			ExchangeConfigJson: "{}",
		}

		marketPrices[i] = pricestypes.MarketPrice{
			Id: uint32(i),
			Price: uint64(1),
			Exponent: -8,
		}
	}

	genKVs = append(genKVs, chain.GenesisKV{
		Key:   "app_state.prices.market_params",
		Value: marketParams,
	})
	genKVs = append(genKVs, chain.GenesisKV{
		Key:   "app_state.prices.market_prices",
		Value: marketPrices,
	})

	return chain.ModifyGenesis(genKVs)
}

func GetProvider(ctx context.Context, logger *zap.Logger) (provider.Provider, error) {
	switch os.Getenv(envProviderType) {
	case digitalOceanProviderType:
		return digitalocean.NewDigitalOceanProvider(
			ctx,
			logger,
			"slinky-digital-ocean",
			os.Getenv(envDigitalOceanAPIKey),
		)
	case dockerProviderType:
		return docker.NewDockerProvider(
			ctx,
			logger,
			"slinky-docker",
		)
	default:
		return nil, fmt.Errorf("unknown provider type: %s", os.Getenv(envProviderType))
	}
}

func GetChain(ctx context.Context, logger *zap.Logger, config petritypes.ChainConfig) (petritypes.ChainI, provider.Provider, error) {
	prov, err := GetProvider(ctx, logger)
	if err != nil {
		return nil, nil, err
	}
	chain, err := chain.CreateChain(
		ctx,
		logger,
		prov,
		config,
	)

	return chain, prov, err
}
