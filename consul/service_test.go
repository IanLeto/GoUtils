package consul

import (
	"GoUtils/utils"
	"github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TestServiceSuit struct {
	suite.Suite
	Service *BaseService
	HTTPOpt *utils.HttpServiceOption
}

func (s *TestServiceSuit) SetupTest() {
	s.HTTPOpt = &utils.HttpServiceOption{
		Addr: "localhost", Port: "9094",
	}
}
func (s *TestServiceSuit) TestBaseService_Check() {
	s.NoError(s.Service.ConsulAgent.ServiceRegister(&api.AgentServiceRegistration{
		ID:      "test",
		Name:    "11111",
		Port:    80,
		Address: "172.17.0.4",
		Tags:    []string{"12"},
		Check: &api.AgentServiceCheck{
			HTTP:     "http://172.17.0.4:80",
			Method:   "GET",
			Interval: (1 * time.Second).String(),
			Timeout:  (10 * time.Second).String(),
		},
	}))
	utils.RunHttpService()
}
func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(TestServiceSuit))
}
