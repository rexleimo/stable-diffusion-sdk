package config

import (
	"fmt"
	"io/ioutil"
	"sync"

	"gopkg.in/yaml.v2"
)

var once sync.Once
var systemConfig *Config

type Config struct {
	SDServer      SDServer      `yaml:"sdserver"`
	MongoDB       MongoDB       `yaml:"mongodb"`
	MiniAppConfig MiniAppConfig `yaml:"miniapp"`
	JwtConfig     JwtConfig     `yaml:"jwt"`
}

func GetConfig() *Config {
	once.Do(func() {
		yamlFile, err := ioutil.ReadFile(".config.yaml")
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		// 读取配置信息
		config := Config{}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			panic(fmt.Errorf("unable to decode into struct, %v", err))
		}
		systemConfig = &config
	})
	return systemConfig
}
