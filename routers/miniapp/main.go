package miniapp

import (
	"stable-diffusion-sdk/core/httpserver"
)

func Init() {
	rg := httpserver.GetInstance().Group("miniapp")

	rg.GET("open_id", getOpenId)
}
