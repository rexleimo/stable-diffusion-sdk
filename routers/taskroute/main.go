package taskroute

import (
	"stable-diffusion-sdk/core/httpserver"
	"stable-diffusion-sdk/middlewares"
)

func Init() {
	rg := httpserver.GetInstance().Group("task")
	rg.Use(middlewares.MiniAppAuthRequired())
	{
		rg.GET("", getTaskList)
		rg.POST("/ids", getTaskListByIds)
	}
}
