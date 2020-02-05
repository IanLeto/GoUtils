package consul

import (
	"GoUtils/utils"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/hashicorp/consul/api"
	"net/http"
	"strings"
	"time"
)

type BaseService struct {
	Name        string
	TTL         time.Duration
	RedisClient redis.UniversalClient
	Cli         *redis.Client
	HttpClient  *http.Client
	ConsulAgent *api.Agent
}



func (s *BaseService) Check() {
	_, err := s.RedisClient.Ping().Result()
	utils.NoErr(err)
}

func (s *BaseService) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	status := 200
	key := strings.Trim(r.URL.Path, "/")
	val, err := s.RedisClient.Get(key).Result()
	if err != nil {
		http.Error(w, "key not found", http.StatusNotFound)
		status = 404
	}
	//go s.update()
	fmt.Println(w, val)
	fmt.Println(status)
}

func (s *BaseService) update() {
	for range time.NewTicker(s.TTL / 2).C {
		utils.NoErr(s.ConsulAgent.PassTTL(s.Name, "still alive"))
	}
}

func (s *BaseService) RegisterServer() {
	serverDef := &api.AgentServiceRegistration{
		Name: s.Name,
		Check: &api.AgentServiceCheck{
			TTL: s.TTL.String(),
		},
	}
	fmt.Println(12)
	utils.NoErr(s.ConsulAgent.ServiceRegister(serverDef))
	go s.update()
}
