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
	//s.Service = &BaseService{
	//	Name: "test",
	//	TTL:  20 * time.Second,
	//	Cli: goredis.NewClient(&goredis.Options{
	//		Addr: "localhost:16379",
	//	}),
	//}
	s.Service = NewBaseService()
	//	s.Service.RegisterServer()
	//	http.HandleFunc("/member", s.Service.ServerHTTP)
	//	http.ListenAndServe(":9999", nil)
}
func (s *TestServiceSuit) TestBaseService_Check() {
	//go utils.NewHttpServer(s.HTTPOpt)
	//p, err := strconv.Atoi(s.HTTPOpt.Port)
	//s.NoError(err)
	s.NoError(s.Service.ConsulAgent.ServiceRegister(&api.AgentServiceRegistration{
		ID:      "test2",
		Name:    "httpServer",
		Port:    80,
		Address: "172.17.0.3",
		Tags:    []string{"12"},
		Check: &api.AgentServiceCheck{
			HTTP:     "http://172.17.0.3:80",
			Method:   "GET",
			Interval: (1 * time.Second).String(),
			Timeout:  (10 * time.Second).String(),
		},
	}))
	utils.RunHttpService()
	//time.Sleep(100 * time.Second)
}

func (s *TestServiceSuit)TestBaseServiceHealth()  {

}

func TestServiceSuite(t *testing.T) {
	suite.Run(t, new(TestServiceSuit))
}
