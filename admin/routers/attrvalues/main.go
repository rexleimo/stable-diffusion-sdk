package attrvalues

import (
	"stable-diffusion-sdk/admin/utils"
)

func Init() {
	rg := utils.GetAdminRouter().Group("attrvalues")

	rg.GET("", list)
	rg.GET("post", posts)
	rg.POST("post", create)
}
