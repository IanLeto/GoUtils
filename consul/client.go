package consul

import (
	"github.com/hashicorp/consul/api"
)

type ConfigCenter struct {
	Client *api.Client
	KV     *api.KV
}

