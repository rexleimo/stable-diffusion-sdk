package qrcode

import (
	"stable-diffusion-sdk/core/httpserver"
	"stable-diffusion-sdk/middlewares"
)

func Init() {
	rg := httpserver.GetInstance().Group("sd")
	rg.Use(middlewares.MiniAppAuthRequired())
	{
		rg.POST("qrcode", progressQrcode)
	}
}
