package text2img

import (
	"stable-diffusion-sdk/core/httpserver"
	"stable-diffusion-sdk/middlewares"
)

func Init() {
	group := httpserver.GetInstance().Group("sd")
	group.Use(middlewares.MiniAppAuthRequired())
	{
		group.POST("/text2img", sendText2ImgTask)
		group.POST("/lighting", lightingText2Img)
	}

}
