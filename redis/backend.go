package redis

import (
	"context"
	goredis "github.com/go-redis/redis"
	"sync"
)

type Backend struct {
	addr string
	cli  *goredis.Client
	pool *sync.Pool
}

func NewBackend(ctx context.Context) *Backend {
	return &Backend{
		addr: "",
		cli: goredis.NewClient(&goredis.Options{
			Addr: "", DB: 1,
		}),
		pool: &sync.Pool{
			New: func() interface{} {
				return make([]string, 0)
			},
		},
	}
}

func NewSentinel(name string) ([]string, error) {
	cli := goredis.NewSentinelClient(&goredis.Options{
		Addr: "",
	})
	return cli.GetMasterAddrByName(name).Result()
}

func (c *Backend) process() {
	defer c.cli.Close()
	c.cli.LPush("key")
	c.cli.RPop("key")
	if err := c.cli.Pipeline().Process(goredis.NewBoolCmd()); err != nil {
		panic(err)
	}
	if err := c.cli.Pipeline().Close(); err != nil {
		panic(err)
	}

}

func NewRedisClient() *goredis.Client {
	return goredis.NewClient(&goredis.Options{
		Addr: "localhost:6379", DB: 0,
	})
}

//func Producer(cli *goredis.Client, killChan <-chan error) {
//	incrment := func(key string) error {
//		txF := func(tx *goredis.Tx) {
//			// 指定key
//			count, err := tx.Get(key).Int()
//			if err != nil {
//				panic(err)
//			}
//			count++
//			_, err = tx.Pipelined(func(pipeline goredis.Pipeliner) error {
//				pipeline.Set(key, count, 0)
//				return nil
//			})
//			if err != nil {
//				panic(err)
//			}
//			return
//		}
//	}
//
//	for retry := 100; retry > 0; retry-- {
//
//	}
//	var err error
//	err = json.Unmarshal(nil, nil)
//
//	if err != nil {
//		panic(err)
//	}
//
//
//}
