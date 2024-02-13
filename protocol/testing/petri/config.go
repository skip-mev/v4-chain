package petri

import (
	"context"
	petritypes "github.com/skip-mev/petri/core/v2/types"
	"github.com/skip-mev/petri/core/v2/provider"
	"github.com/skip-mev/petri/core/v2/provider/docker"
	"github.com/skip-mev/petri/cosmos/v2/chain"
	"github.com/skip-mev/petri/cosmos/v2/node"
	"github.com/dydxprotocol/v4-chain/protocol/testutil/encoding"
	"go.uber.org/zap"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
)

func GetChainConfig() petritypes.ChainConfig {
	return petritypes.ChainConfig{
		Denom:         "stake",
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
			Image: "skip-mev/slinky-e2e-oracle",
			UID:   "1000",
			GID:   "1000",
		},
		GasPrices:      "0stake",
		GasAdjustment:  1.5,
		Bech32Prefix:   "dydx",
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
			Bech32Prefix:     "cosmos",
			HDPath:           hd.CreateHDPath(0, 0, 0),
			SigningAlgorithm: "secp256k1",
		},
		UseGenesisSubCommand: true,
		NodeCreator:          node.CreateNode,
	}
}

func GetGenesisModifier() petritypes.GenesisModifier {
	return chain.ModifyGenesis(nil)
}

func GetProvider(ctx context.Context, logger *zap.Logger) (provider.Provider, error) {
	return docker.NewDockerProvider(
		ctx,
		logger,
		"slinky-docker",
	)
}

func GetChain(ctx context.Context, logger *zap.Logger) (petritypes.ChainI, error) {
	prov, err := GetProvider(ctx, logger)
	if err != nil {
		return nil, err
	}
	return chain.CreateChain(
		ctx,
		logger,
		prov,
		GetChainConfig(),
	)
}
