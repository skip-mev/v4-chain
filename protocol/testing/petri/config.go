package petri

import (
	"context"
	"math/big"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/dydxprotocol/v4-chain/protocol/testutil/encoding"
	"github.com/skip-mev/petri/core/v2/provider"
	"github.com/skip-mev/petri/core/v2/provider/docker"
	petritypes "github.com/skip-mev/petri/core/v2/types"
	"github.com/skip-mev/petri/cosmos/v2/chain"
	"github.com/skip-mev/petri/cosmos/v2/node"
	"go.uber.org/zap"
)

const (
	denom =  "dv4tnt"
	prefix = "dydx"
) 

func GetChainConfig() petritypes.ChainConfig {
	return petritypes.ChainConfig{
		Denom:         denom,
		Decimals:      6,
		NumValidators: 4,
		NumNodes:      2,
		BinaryName:    "dydxprotocold",
		Image: provider.ImageDefinition{
			Image: "dydxprotocol-base",
			UID:   "1000",
			GID:   "1000",
		},
		SidecarImage: provider.ImageDefinition{
			Image: "dydxprotocol-base",
			UID:   "1000",
			GID:   "1000",
		},
		SidecarArgs: []string{"slinky", "--oracle-config-path", "/etc/oracle.toml", "-host", "0.0.0.0", "-port", "8080"},
		GasPrices:      "0dv4tnt",
		GasAdjustment:  1.5,
		Bech32Prefix:   prefix,
		EncodingConfig: encoding.GetTestEncodingCfg(),
		HomeDir:        "/petri-test",
		SidecarHomeDir: "/petri-test",
		SidecarPorts:   []string{"8080"},
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
		NodeCreator:          node.CreateNode,
		GenesisDelegation: big.NewInt(10_000_000_000_000),
		GenesisBalance: big.NewInt(100_000_000_000_000),
	}
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
	return chain.ModifyGenesis(genKVs)
}

func GetProvider(ctx context.Context, logger *zap.Logger) (provider.Provider, error) {
	return docker.NewDockerProvider(
		ctx,
		logger,
		"slinky-docker",
	)
}

func GetChain(ctx context.Context, logger *zap.Logger, config petritypes.ChainConfig) (petritypes.ChainI, error) {
	prov, err := GetProvider(ctx, logger)
	if err != nil {
		return nil, err
	}
	return chain.CreateChain(
		ctx,
		logger,
		prov,
		config,
	)
}
