package routers

import (
	attrs "stable-diffusion-sdk/routers/attr"
	"stable-diffusion-sdk/routers/attrvalues"
	"stable-diffusion-sdk/routers/category"
	"stable-diffusion-sdk/routers/miniapp"
	"stable-diffusion-sdk/routers/sd-router/avatar"
	img2img "stable-diffusion-sdk/routers/sd-router/img-2-img"
	"stable-diffusion-sdk/routers/sd-router/progress"
	"stable-diffusion-sdk/routers/sd-router/qrcode"
	text2img "stable-diffusion-sdk/routers/sd-router/text-2-img"
	"stable-diffusion-sdk/routers/stylesroute"
	"stable-diffusion-sdk/routers/taskroute"
	"stable-diffusion-sdk/routers/userroute"
)

func Init() {
	text2img.Init()
	img2img.Init()
	qrcode.Init()
	progress.Init()
	category.Init()
	attrs.Init()
	attrvalues.Init()
	miniapp.Init()
	taskroute.Init()
	userroute.Init()
	stylesroute.Init()
	avatar.Init()
}
