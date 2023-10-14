package routers

import (
	attrs "stable-diffusion-sdk/routers/attr"
	"stable-diffusion-sdk/routers/attrvalues"
	"stable-diffusion-sdk/routers/category"
	"stable-diffusion-sdk/routers/miniapp"
	"stable-diffusion-sdk/routers/sd-router/progress"
	text2img "stable-diffusion-sdk/routers/sd-router/text-2-img"
	"stable-diffusion-sdk/routers/taskroute"
	"stable-diffusion-sdk/routers/userroute"
)

func Init() {
	text2img.Init()
	progress.Init()
	category.Init()
	attrs.Init()
	attrvalues.Init()
	miniapp.Init()
	taskroute.Init()
	userroute.Init()
}
