package petri

import (
	tmloadtest "github.com/informalsystems/tm-load-test/pkg/loadtest"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	petritypes "github.com/skip-mev/petri/core/v2/types"
	"github.com/skip-mev/petri/cosmos/v2/cosmosutil"
	"github.com/skip-mev/petri/cosmos/v2/loadtest"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	assetsmodule "github.com/dydxprotocol/v4-chain/protocol/x/assets"
	blocktimemodule "github.com/dydxprotocol/v4-chain/protocol/x/blocktime"
	bridgemodule "github.com/dydxprotocol/v4-chain/protocol/x/bridge"
	clobmodule "github.com/dydxprotocol/v4-chain/protocol/x/clob"
	delaymsgmodule "github.com/dydxprotocol/v4-chain/protocol/x/delaymsg"
	epochsmodule "github.com/dydxprotocol/v4-chain/protocol/x/epochs"
	feetiersmodule "github.com/dydxprotocol/v4-chain/protocol/x/feetiers"
	govplusmodule "github.com/dydxprotocol/v4-chain/protocol/x/govplus"
	perpetualsmodule "github.com/dydxprotocol/v4-chain/protocol/x/perpetuals"
	pricesmodule "github.com/dydxprotocol/v4-chain/protocol/x/prices"
	ratelimitmodule "github.com/dydxprotocol/v4-chain/protocol/x/ratelimit"
	rewardsmodule "github.com/dydxprotocol/v4-chain/protocol/x/rewards"
	sendingmodule "github.com/dydxprotocol/v4-chain/protocol/x/sending"
	statsmodule "github.com/dydxprotocol/v4-chain/protocol/x/stats"
	subaccountsmodule "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts"
	vestmodule "github.com/dydxprotocol/v4-chain/protocol/x/vest"
	evidencemodule "cosmossdk.io/x/evidence"
	feegrantmodule "cosmossdk.io/x/feegrant/module"
	"cosmossdk.io/x/upgrade"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/consensus"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/ibc-go/modules/capability"
	ica "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v8/modules/core"
	custommodule "github.com/dydxprotocol/v4-chain/protocol/app/module"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	envKeepAlive = "PETRI_ENV_KEEP_ALIVE"
)

// SlinkyIntegrationSuite is a test-suite used to spin up load-tests of arbitrary size for dydx nodes
type SlinkyIntegrationSuite struct {
	suite.Suite

	logger *zap.Logger

	spec *petritypes.ChainConfig

	chain petritypes.ChainI
}

func NewSlinkyIntegrationSuite(spec *petritypes.ChainConfig) *SlinkyIntegrationSuite {
	return &SlinkyIntegrationSuite{
		spec: spec,
	}
}

func (s *SlinkyIntegrationSuite) SetupSuite() {
	// create the logger
	var err error
	s.logger, err = zap.NewDevelopment()
	s.Require().NoError(err)

	// create the chain
	s.chain, err = GetChain(context.Background(), s.logger)
	s.Require().NoError(err)

	//initialize the chain
	err = s.chain.Init(context.Background())
	s.Require().NoError(err)
}

func (s *SlinkyIntegrationSuite) TearDownSuite() {
	// get the oracle integration-test suite keep alive env
	if ok := os.Getenv(envKeepAlive); ok == "" {
		return
	}
	err := s.chain.Teardown(context.Background())
	s.Require().NoError(err)
}

func (s *SlinkyIntegrationSuite) TestSlinkyUnderLoad() {
	err := s.chain.WaitForHeight(context.Background(), 1)
	s.Require().NoError(err)

	encCfg := cosmosutil.EncodingConfig{
		InterfaceRegistry: s.chain.GetInterfaceRegistry(),
		Codec:             codec.NewProtoCodec(s.chain.GetInterfaceRegistry()),
		TxConfig:          s.chain.GetTxConfig(),
	}

	clientFactory, err := loadtest.NewDefaultClientFactory(
		loadtest.ClientFactoryConfig{
			Chain:                 s.chain,
			Seeder:                cosmosutil.NewInteractingWallet(s.chain, s.chain.GetFaucetWallet(), encCfg),
			WalletConfig:          s.spec.WalletConfig,
			AmountToSend:          1000000000,
			SkipSequenceIncrement: false,
			EncodingConfig:        encCfg,
			MsgGenerator:          s.genMsg,
		},
		getModuleBasics(),
	)
	s.Require().NoError(err)

	err = tmloadtest.RegisterClientFactory("slinky", clientFactory)
	s.Require().NoError(err)

	var endpoints []string
	for _, val := range s.chain.GetValidators() {
		endpoint, err := val.GetTMClient(context.Background())
		s.Require().NoError(err)

		url := strings.Replace(endpoint.Remote(), "http", "ws", -1)

		endpoints = append(endpoints, fmt.Sprintf("%s/websocket", url))
	}

	cfg := tmloadtest.Config{
		ClientFactory:        "slinky",
		Connections:          1,
		Endpoints:            endpoints,
		Time:                 60,
		SendPeriod:           1,
		Rate:                 350,
		Size:                 250,
		Count:                -1,
		BroadcastTxMethod:    "async",
		EndpointSelectMethod: "supplied",
	}
	err = tmloadtest.ExecuteStandalone(cfg)
	s.Require().NoError(err)
}

func getModuleBasics() module.BasicManager {
	return module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.NewAppModuleBasic(genutiltypes.DefaultMessageValidator),
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			[]govclient.ProposalHandler{
				paramsclient.ProposalHandler,
			},
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		custommodule.SlashingModuleBasic{},
		evidencemodule.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		ibc.AppModuleBasic{},
		ica.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		transfer.AppModuleBasic{},
		consensus.AppModuleBasic{},
		authzmodule.AppModuleBasic{},

		// Custom modules
		pricesmodule.AppModuleBasic{},
		assetsmodule.AppModuleBasic{},
		blocktimemodule.AppModuleBasic{},
		bridgemodule.AppModuleBasic{},
		feetiersmodule.AppModuleBasic{},
		perpetualsmodule.AppModuleBasic{},
		statsmodule.AppModuleBasic{},
		subaccountsmodule.AppModuleBasic{},
		clobmodule.AppModuleBasic{},
		vestmodule.AppModuleBasic{},
		rewardsmodule.AppModuleBasic{},
		sendingmodule.AppModuleBasic{},
		govplusmodule.AppModuleBasic{},
		delaymsgmodule.AppModuleBasic{},
		epochsmodule.AppModuleBasic{},
		ratelimitmodule.AppModuleBasic{},
	)
}

func (s *SlinkyIntegrationSuite) genMsg(senderAddress []byte) ([]sdk.Msg, petritypes.GasSettings, error) {
	return []sdk.Msg{
			&banktypes.MsgSend{
				FromAddress: string(senderAddress),
				ToAddress:   "cosmos1qy3523p8x9z0j6z3qg3y7t4v6gj6z9q8r9m9x5",
				Amount:      sdk.NewCoins(sdk.NewInt64Coin("stake", 1)),
			},
		}, petritypes.GasSettings{
			Gas:         200000,
			GasDenom:    "stake",
			PricePerGas: 0,
		}, nil
}