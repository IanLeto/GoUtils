package Counterexample

import (
	"encoding/json"
	"io/ioutil"
)

type RedisConfig struct {
	Url  string
	Port int
}

func NewRedisConifg() *RedisConfig {
	return &RedisConfig{}
}

func GetNewRedisConfig(path string) *RedisConfig {
	var conf RedisConfig
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &conf); err != nil {
		panic(err)
	}
	return &conf
}
