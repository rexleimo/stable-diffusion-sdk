package http

import (
	"stable-diffusion-sdk/utils/config"
	"sync"

	"github.com/go-resty/resty/v2"
)

var once sync.Once

var instance *resty.Client

func GetInstance() *resty.Client {
	once.Do(func() {
		instance = resty.New()
	})
	return instance
}

func GetSDServer() *resty.Request {
	return GetInstance().SetBaseURL(config.GetConfig().SDServer.Host).R()
}
