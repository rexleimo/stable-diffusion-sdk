package routers

import (
	"stable-diffusion-sdk/admin/routers/attrs"
	"stable-diffusion-sdk/admin/routers/attrvalues"
	"stable-diffusion-sdk/admin/routers/category"
)

func Init() {
	attrs.Init()
	category.Init()
	attrvalues.Init()
}
