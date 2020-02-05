package consul

import "transfer/consul"

type KVAPI interface {
	Get(key string, q *consul.QueryOptions) (*consul.KVPair, *consul.QueryMeta, error)
}
