package attrs

import "stable-diffusion-sdk/admin/utils"

func Init() {
	rg := utils.GetAdminRouter().Group("attrs")
	rg.GET("", getList)
	rg.GET("post", post)
	rg.POST("post", create)
	rg.PUT(":id", update)
	rg.DELETE(":id", delete)
}
