package consul

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"transfer/consul"
)

type TestinterfaceInTransferSuite struct {
	suite.Suite
	cli consul.Client
}

func (s *TestinterfaceInTransferSuite) TestFunc() {

}

func (s *TestinterfaceInTransferSuite) SetupTest() {

}

func TestInterfaceConsulSuite(t *testing.T) {
	suite.Run(t, new(TestinterfaceInTransferSuite))
}
