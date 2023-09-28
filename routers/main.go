package routers

import (
	"stable-diffusion-sdk/core/httpserver"
	text2img "stable-diffusion-sdk/routers/sd-router/text-2-img"
)

func Init() {
	text2img.Init()
	httpserver.GetInstance().Run(":7100")
}
