package styles

import "stable-diffusion-sdk/admin/utils"

func Init() {
	rg := utils.GetAdminRouter().Group("styles")
	{
		rg.GET("", renderList)
		rg.GET("/posts", renderForm)
		rg.GET("/edit/:id", renderEditForm)
		rg.POST("/posts", postForm)
		rg.PUT("/edit/:id", editForm)
		rg.DELETE(":id", deleteForm)
	}
}
