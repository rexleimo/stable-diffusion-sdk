package userroute

import (
	"stable-diffusion-sdk/core/httpserver"
	"stable-diffusion-sdk/middlewares"
)

func Init() {
	rg := httpserver.GetInstance().Group("user")
	rg.Use(middlewares.MiniAppAuthRequired())
	{
		rg.GET("check_in", checkIn)
		rg.GET("info", userInfo)
		rg.POST("replenish", replenish)
	}
}
