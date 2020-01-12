package consul

import (
	"GoUtils/utils"
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

type ConfigCenter struct {
	Client *api.Client
	KV     *api.KV
}

func NewConfig() *api.Config {
	var (
		conf = viper.New()
	)
	conf.SetConfigName("consulConfig")
	conf.AddConfigPath("/workDir/Go/goPath/src/GoUtils")
	conf.SetConfigType("yaml")
	utils.NoErr(conf.ReadInConfig())
	config := api.DefaultConfig()
	config.Address = conf.GetString("Addr") + ":" + conf.GetString("Port")
	return config
}

func NewConsulClient(ctx context.Context) *api.Client {
	cli, err := api.NewClient(NewConfig())
	utils.NoErr(err)
	return cli
}

func NewConsulConfig(ctx context.Context) *ConfigCenter {
	var (
		conf = viper.New()
	)
	conf.SetConfigName("consulConfig")
	conf.AddConfigPath("/workDir/Go/goPath/src/GoUtils")
	conf.SetConfigType("yaml")
	utils.NoErr(conf.ReadInConfig())
	config := api.DefaultConfig()
	config.Address = conf.GetString("Addr") + ":" + conf.GetString("Port")
	fmt.Println(config.Address)
	cli, err := api.NewClient(config)
	utils.NoErr(err)

	kv := cli.KV()

	return &ConfigCenter{
		Client: cli,
		KV:     kv,
	}
}

func RegesisterService() {
	//service := NewConsulConfig(nil)
	//err := service.Client.Agent().ServiceRegister(&api.AgentServiceRegistration{
	//
	//})
}

func DemoLock() {

	//cli := NewConsulConfig(nil)
	//lockKey := "demo-lock-key"
	//lock, err := cli.Client.LockOpts(&api.LockOptions{
	//	Key:         lockKey,
	//	Value:       []byte("sender 1"),
	//	SessionTTL:  "10s",
	//	SessionName: "xx",
	//})

}
