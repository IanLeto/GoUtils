package consul

import (
	consulAPI "github.com/hashicorp/consul/api"
)

type ClientAPI interface {
	Client() *consulAPI.Client
	KV() KvAPI
}
type ClientConsul struct{}

func (ClientConsul) Client() *consulAPI.Client {
	panic("implement me")
}

func (ClientConsul) KV() KvAPI {
	panic("implement me")
}

type KvAPI interface {
	Get()
}


