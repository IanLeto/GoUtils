package consul

import (
	"GoUtils/utils"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestConsulClientSuite struct {
	suite.Suite
	Cli *api.Client
}

func (s *TestConsulClientSuite) SetupTest() {
	s.Cli = utils.NewConsulClient()
}

func (s *TestConsulClientSuite) TestKVAPI() {
	kv := s.Cli.KV()
	_, err := s.Cli.KV().Put(&api.KVPair{
		Key:   "hello",
		Value: []byte(`world`),
	}, &api.WriteOptions{})
	value, _, err := s.Cli.KV().Get("hello", nil)
	s.NoError(err)
	s.Equal("world", string(value.Value))
	res, _, err := kv.Keys("", "", nil)
	s.Equal(2, len(res))
	_, err = s.Cli.KV().Delete("hello", nil)
	s.NoError(err)
}
func (s *TestConsulClientSuite) TestShowMetaQuery() {
	kv := s.Cli.KV()
	writeMeta, err := kv.Put(&api.KVPair{
		Key:         "hello",
		CreateIndex: 10086,
		ModifyIndex: 32,
		//LockIndex:
		Flags: 1,
		Value: []byte(`world`),
		//Session:
	}, &api.WriteOptions{
		Datacenter: "dc1",
		//Token:
	})
	s.NoError(err)
	fmt.Println(writeMeta.RequestTime)

}
func (s *TestConsulClientSuite) TestKVPut() {

}

func (s *TestConsulClientSuite) TestKVService() {
}

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(TestConsulClientSuite))
}
