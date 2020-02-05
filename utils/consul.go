package utils

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

func NewConfig() *api.Config {
	var (
		conf = viper.New()
	)
	conf.SetConfigName("consulConfig")
	conf.AddConfigPath("/workDir/Go/goPath/src/GoUtils")
	conf.SetConfigType("yaml")
	NoErr(conf.ReadInConfig())
	config := api.DefaultConfig()
	config.Address = conf.GetString("Addr") + ":" + conf.GetString("Port")
	return config
}

func NewConsulClient() *api.Client {
	var (
		conf = viper.New()
	)
	conf.SetConfigName("consulConfig")
	conf.AddConfigPath("/workDir/Go/goPath/src/GoUtils")
	conf.SetConfigType("yaml")
	NoErr(conf.ReadInConfig())
	config := api.DefaultConfig()
	config.Address = conf.GetString("Addr") + ":" + conf.GetString("Port")
	fmt.Println(config.Address)
	cli, err := api.NewClient(config)
	NoErr(err)
	return cli

}
