package consul

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestNewConsul(t *testing.T) {

}

type TestConsulSuite struct {
	suite.Suite
}

func (s *TestConsulSuite) TestConnect() {
	NewConsul()

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestConsulSuite))
}

