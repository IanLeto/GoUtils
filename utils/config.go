package utils

import "github.com/spf13/viper"

type ConsulConfig struct {
}

func ReadConfig(cv *viper.Viper, name, path, t string) {
	//cv := viper.New()
	cv.SetConfigFile(name)
	cv.AddConfigPath(path)
	cv.SetConfigType(t)
}
