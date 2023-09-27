package routers

import (
	httpserver "stable-diffusion-sdk/core/http-server"
	text2img "stable-diffusion-sdk/routers/sd-router/text-2-img"
)

func Init() {
	text2img.Init()
	httpserver.GetInstance().Run(":7100")
}
