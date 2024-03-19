package petri

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/dydxprotocol/v4-chain/protocol/testutil/encoding"
	pricestypes "github.com/dydxprotocol/v4-chain/protocol/x/prices/types"
	subaccounttypes "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"
	"github.com/skip-mev/petri/core/v2/provider"
	"github.com/skip-mev/petri/core/v2/provider/digitalocean"
	"github.com/skip-mev/petri/core/v2/provider/docker"
	petritypes "github.com/skip-mev/petri/core/v2/types"
	"github.com/skip-mev/petri/cosmos/v2/chain"
	"github.com/skip-mev/petri/cosmos/v2/node"
	oracleconfig "github.com/skip-mev/slinky/oracle/config"
	mmtypes "github.com/skip-mev/slinky/x/marketmap/types"
	"go.uber.org/zap"
)

const (
	denom            = "dv4tnt"
	usdcDenom 	  = "ibc/8E27BA2D5493AF5636760E354E46004562C46AB7EC0CC4C1CA14E9E20E2545B5"
	prefix           = "dydx"
	homeDir          = "/petri-test"
	oracleHomeDir = "/oracle"
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
	faucetAccount = "dydx1nzuttarf5k2j0nug5yzhr6p74t9avehn9hlh8m"
	url = "https://api.gateio.ws/api/v4/spot/currency_pairs"
	genesisUSDCAmount = 100000000000000000
)
var ( 
	doRegions = []string{"nyc1", "nyc3"}
	cdc codec.Codec
)

func init() {
	ir := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)
	cdc = codec.NewProtoCodec(ir)
}


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
			Image: "nikhilv01/dydxprotocol-base:latest",
			UID:   "1000",
			GID:   "1000",
		},
		SidecarImage: provider.ImageDefinition{
			Image: "nikhilv01/dydxprotocol-base:latest",
			UID:   "1000",
			GID:   "1000",
		},
		SidecarArgs:    []string{
			"slinky", 
			"--oracle-config-path", fmt.Sprintf("%s/%s", oracleHomeDir, oracleConfigPath), 
			"--market-config-path", fmt.Sprintf("%s/%s", oracleHomeDir, marketConfigPath), 
			"-host", "0.0.0.0", 
			"-port", "8080",
		},
		GasPrices:      "0dv4tnt",
		GasAdjustment:  1.5,
		Bech32Prefix:   prefix,
		EncodingConfig: encoding.GetTestEncodingCfg(),
		HomeDir:        homeDir,
		SidecarHomeDir: oracleHomeDir,
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
		NodeCreator: func(ctx context.Context, l *zap.Logger, nc petritypes.NodeConfig) (petritypes.NodeI, error) {
			nodeI, err := node.CreateNode(ctx, l, nc)
			if err != nil {
				return nil, err
			}

			n, ok := nodeI.(*node.Node)
			if !ok {
				return nil, fmt.Errorf("node is expected to be of type: %v, but is %v", (*node.Node)(nil), n)
			}

			if len(n.Sidecars) != 1 {
				return nil, fmt.Errorf("node has %d sidecars instead of 1 oracle sidecar", len(n.Sidecars))
			}

			n.Sidecars[0].PreStart = oraclePreStart
			return n, nil
		}, // modify to account for additional parameters
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

// oraclePreStart writes the default oracle configs to the /oracle dir in the sidecar container
func oraclePreStart(ctx context.Context, oracle *provider.Task) error {
	oracle.Logger().Info("writing oracle config in pre-start", zap.String("sidecar", oracle.Definition.Name))

	// if the config already exists, do not overwrite it
	if _, err := oracle.ReadFile(ctx, oracleConfigPath); err == nil {
		oracle.Logger().Info("oracle config already exists, skipping write", zap.String("sidecar", oracle.Definition.Name))
		return nil
	}

	cfg := oracleconfig.OracleConfig{
		UpdateInterval: 1 * time.Second,
		MaxPriceAge:   1 * time.Second,	
	}
	bz, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	// write oracle config
	if err := oracle.WriteFile(ctx, oracleConfigPath, bz); err != nil {
		return err
	}

	mm := mmtypes.MarketMap{}

	bz, err = json.Marshal(mm)
	if err != nil {
		return err
	}

	// write market config
	return oracle.WriteFile(ctx, marketConfigPath, bz)
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

	// setup subaccounts
	accounts := []string{faucetAccount}
	subaccounts := make([]subaccounttypes.Subaccount, len(accounts))
	for i, addr := range accounts {
		subaccounts[i] = subaccounttypes.Subaccount{
			Id: &subaccounttypes.SubaccountId{
				Owner: addr,
				Number: 0,
			},
			MarginEnabled: true,
		}
	}
	genKVs = append(genKVs, chain.GenesisKV{
		Key:   "app_state.subaccounts.subaccounts",
		Value: subaccounts,
	})

	return func(b []byte) ([]byte, error) {
		genBz, err := chain.ModifyGenesis(genKVs)(b)
		if err != nil {
			return nil, err
		}

		// unmarshal genBz and update account states
		var genState map[string]json.RawMessage
		if err := json.Unmarshal(genBz, &genState); err != nil {
			return nil, err
		}

		appStateBz := genState["app_state"]
		var appState map[string]json.RawMessage
		if err := json.Unmarshal(appStateBz, &appState); err != nil {
			return nil, err
		}

		// update account states
		updatedAppState, err := updateGenesisAccounts(appState, accounts)
		if err != nil {
			return nil, err
		}

		genState["app_state"], err = json.Marshal(updatedAppState)
		if err != nil {
			return nil, err
		}

		return json.Marshal(genState)
	}
}

type JSONAccountState struct {
	Type string `json:"@type"`
	Address string `json:"address"`
	Pubkey []byte `json:"pub_key"`
	AccountNumber uint64 `json:"account_number"`
	Sequence uint64 `json:"sequence"`
}

func updateGenesisAccounts(genesis map[string]json.RawMessage, accounts []string) (map[string]json.RawMessage, error) {
	// setup in auth state
	genState, err := setupAuthState(genesis, accounts)
	if err != nil {
		return nil, err
	}

	// setup in bank state
	genState, err = setupBankState(genState, accounts)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}

func setupBankState(genesis map[string]json.RawMessage, accounts []string) (map[string]json.RawMessage, error) {
	bankBz, ok := genesis[banktypes.ModuleName]
	if !ok {
		return nil, fmt.Errorf("bank module not found in genesis")
	}

	var bankGenesis banktypes.GenesisState
	if err := cdc.UnmarshalJSON(bankBz, &bankGenesis); err != nil {
		return nil, err
	}

	// setup balances
	balances := make([]banktypes.Balance, len(accounts))
	usdcBalance := sdk.NewCoin(usdcDenom, math.NewInt(100000000000000000))
	dydxAccountBalance := math.NewInt(1000000).Mul(math.NewInt(int64(1e18)))
	dydxBalance := sdk.NewCoin(denom, dydxAccountBalance)
	for i, addr := range accounts {
		balances[i] = banktypes.Balance{
			Address: addr,
			Coins:   sdk.NewCoins(
				usdcBalance,
				dydxBalance,
			),
		}
	}
	bankGenesis.Balances = append(bankGenesis.Balances, balances...)

	genesis[banktypes.ModuleName] = cdc.MustMarshalJSON(&bankGenesis)
	return genesis, nil
}

func setupAuthState(genesis map[string]json.RawMessage, accounts []string) (map[string]json.RawMessage, error) {
	authBz, ok := genesis[authtypes.ModuleName]
	if !ok {
		return nil, fmt.Errorf("auth module not found in genesis")
	}

	var authGenesis authtypes.GenesisState
	if err := cdc.UnmarshalJSON(authBz, &authGenesis); err != nil {
		return nil, err
	}

	// setup accounts
	genAccounts, err := authtypes.UnpackAccounts(authGenesis.Accounts)
	if err != nil {
		return nil, err
	}

	for _, addr := range accounts {
		accAddr, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, err
		}
		bacc := authtypes.NewBaseAccountWithAddress(accAddr)
		genAccounts = append(genAccounts, bacc)
	}

	authGenesis.Accounts, err = authtypes.PackAccounts(genAccounts)
	if err != nil {
		return nil, err
	}

	authBz, err = cdc.MarshalJSON(&authGenesis)
	if err != nil {
		return nil, err
	}
	
	genesis[authtypes.ModuleName] = authBz
	return genesis, nil
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
