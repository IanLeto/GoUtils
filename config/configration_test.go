package config

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TestConsulClientAPISuite struct {
	suite.Suite
	cli *ConsulClient
}

//
func (s *TestConsulClientAPISuite) TestGetFunc() {
	// 注意 key 必须完全正确
	v, _, err := s.cli.ConsulClient.KV().Get("ian/hello/world", nil)
	if err != nil {
		panic(err)
	}
	if v == nil {
		fmt.Println("err")
	} else {
		fmt.Println(v.Value)

	}
}

// 获取prefix
func (s *TestConsulClientAPISuite) TestKeysFunc() {
	// 监听前缀，获取相应的key
	v, _, err := s.cli.ConsulClient.KV().Keys("ian", "", nil)
	if err != nil {
		panic(err)
	}
	if v == nil {
		fmt.Println("err")
	} else {
		for _, value := range v {
			fmt.Println(value)
		}

	}
}

func (s *TestConsulClientAPISuite) SetupTest() {
	var err error
	s.cli, err = NewConsulConfig("config/config", nil)
	if err != nil {
		panic(err)
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TestConsulClientAPISuite))
}
