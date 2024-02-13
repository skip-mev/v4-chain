package petri_test

import (
	"testing"
	"github.com/dydxprotocol/v4-chain/protocol/testing/petri"
	"github.com/stretchr/testify/suite"
)

// runs the slinky integration tests
func TestSlinkyIntegration(t *testing.T) {
	chainCfg := petri.GetChainConfig()
	suite.Run(t, petri.NewSlinkyIntegrationSuite(&chainCfg))
}
