package currencyInGo

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestctxDemo_testSuite struct {
	suite.Suite
}

func (s *TestctxDemo_testSuite) TestFunc() {
	GrandFather()
}

func (s *TestctxDemo_testSuite) SetupTest() {

}

func TestCtxDemoSuite(t *testing.T) {
	suite.Run(t, new(TestctxDemo_testSuite))
}
