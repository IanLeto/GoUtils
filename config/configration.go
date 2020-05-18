package config

import (
	consulApi "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

type Configuration interface {
	Get(key string) interface{}
	// 以class 读取属性的方式，获取配置文件
	GetString(key string) string
	// 返回配置文件路径
	GetPath(key string) string
}

func InitConifgPath(path string) {

}

//
type ConsulClient struct {
	*viper.Viper
	ConsulClient consulApi.Client
}

func NewConsulConfig(path string, conf *consulApi.Config) (*ConsulClient, error) {
	var (
		client *consulApi.Client
		err    error
	)
	viper.AddConfigPath(path)
	if conf == nil {
		client, err = consulApi.NewClient(&consulApi.Config{
			Address: viper.GetString("Address"),
		})
	} else {
		client, err = consulApi.NewClient(conf)
	}

	if err != nil {
		panic(err)
	}
	return &ConsulClient{ConsulClient: *client}, err
}

func (ConsulClient) Get(key string) interface{} {
	panic("implement me")
}

func (ConsulClient) GetString(key string) string {
	panic("implement me")
}

func (ConsulClient) GetPath(key string) string {
	panic("implement me")
}
