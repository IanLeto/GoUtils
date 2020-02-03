package consul

import (
	"github.com/hashicorp/consul/api"
)

type LeaderInterface interface {
	GetAgentName() string
	GetKey(string) (*api.KVPair, error)
	PutKey(pair *api.KVPair) (bool, error)
	ReleaseKey(pair *api.KVPair) (bool, error)
	GetSession(string) string
	AcquireSessionKey(string, string) (bool, error)
	GetHealthChecks(state string, opt *api.QueryOptions) ([]*api.HealthCheck, error)
}

type LeaderElection struct {
	LeaderKey    string
	WatchWaiTime int
	StopElection chan bool
	Client       LeaderInterface
}

func (le *LeaderElection) CancelElection() {
	le.StopElection <- true
}

type LeaderClient struct {
	Client *api.Client
}

func (*LeaderClient) GetAgentName() string {
	panic("implement me")
}

func (*LeaderClient) GetKey(string) (*api.KVPair, error) {
	panic("implement me")
}

func (*LeaderClient) PutKey(pair *api.KVPair) (bool, error) {
	panic("implement me")
}

func (*LeaderClient) ReleaseKey(pair *api.KVPair) (bool, error) {
	panic("implement me")
}

func (*LeaderClient) GetSession(string) string {
	panic("implement me")
}

func (*LeaderClient) AcquireSessionKey(string, string) (bool, error) {
	panic("implement me")
}

func (*LeaderClient) GetHealthChecks(state string, opt *api.QueryOptions) ([]*api.HealthCheck, error) {
	panic("implement me")
}
