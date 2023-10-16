package qrcode

import (
	"stable-diffusion-sdk/core/httpserver"
)

func Init() {
	rg := httpserver.GetInstance().Group("sd")
	// rg.Use(middlewares.MiniAppAuthRequired())
	{
		rg.POST("qrcode", progressQrcode)
	}
}
