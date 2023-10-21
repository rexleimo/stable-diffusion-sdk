package stylesroute

import "stable-diffusion-sdk/core/httpserver"

func Init() {
	rg := httpserver.GetInstance().Group("styles")
	{
		rg.GET("/list", getStyleList)
	}
}
