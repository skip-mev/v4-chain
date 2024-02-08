package petri

import (
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	petritypes "github.com/skip-mev/petri/types/v2"
)

// SlinkyIntegrationSuite is a test-suite used to spin up load-tests of arbitrary size for dydx nodes
type SlinkyIntegrationSuite struct {
	suite.Suite

	logger *zap.Logger

	spec petritypes.ChainConfig
}
