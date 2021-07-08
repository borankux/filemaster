package config

import (
	"github.com/spf13/viper"
	"log"
)

var config *viper.Viper

func Init(env string) {
	config = viper.GetViper()
	config.SetConfigType("yml")
	config.SetConfigName(env)
	config.AddConfigPath("server/config/")
	err := config.ReadInConfig()

	if err != nil {
		log.Fatalf("Failed to load config %s :%v", env, err)
	}
}

func GetConf() *viper.Viper{
	return config
}