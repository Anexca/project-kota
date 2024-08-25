package config

import (
	"github.com/spf13/viper"
)

type Environment struct {
	ServerPort    string `viper:"PORT"`
	RedisPort     string `viper:"REDIS_PORT"`
	RedisAddress  string `viper:"REDIS_ADDRESS"`
	RedisPassword string `viper:"REDIS_PASSWORD"`
	RedisDatabase string `viper:"REDIS_DATABASE"`
}

func LoadEnvironment() (*Environment, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err

	}

	var config Environment
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
