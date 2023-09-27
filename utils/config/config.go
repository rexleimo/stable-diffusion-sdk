package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once
var systemConfig *Config

type Config struct {
	SDServer SDServer `yaml:"sdserver"`
}

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile(".config.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// 读取配置信息
		config := Config{}
		err = viper.Unmarshal(&config)
		if err != nil {
			panic(fmt.Errorf("unable to decode into struct, %v", err))
		}
		systemConfig = &config
	})
	return systemConfig
}
