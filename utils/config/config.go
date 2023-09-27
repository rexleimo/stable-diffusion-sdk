package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

type Config struct {
	SDServer SDServer `yaml:"sdserver"`
}

func GetConfig() *Config {
	viper.SetConfigFile(".config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 读取配置信息
	systemConfig := Config{}
	err = viper.Unmarshal(&systemConfig)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	return &systemConfig
}
