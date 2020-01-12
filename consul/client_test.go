package consul

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestConsulClientSuite struct {
	suite.Suite
	Cli *ConfigCenter
}

func (s *TestConsulClientSuite) SetupTest() {
	s.Cli = NewConsulConfig(nil)
}

func (s *TestConsulClientSuite) TestKVGet() {
	value, _, err := s.Cli.KV.Get("ian/hello/world", nil)
	s.NoError(err)
	s.NotNil(value)
	s.Equal("hello consul", string(value.Value))
}

func (s *TestConsulClientSuite) TestKVPut() {

}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(TestConsulClientSuite))
}
