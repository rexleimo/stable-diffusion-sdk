package category

import (
	"stable-diffusion-sdk/admin/utils"
)

func Init() {
	group := utils.GetAdminRouter().Group("category")
	group.GET("", list)
	group.POST("", create)
	group.PUT(":id", update)
}
