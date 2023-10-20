package styles

import "stable-diffusion-sdk/admin/utils"

func Init() {
	rg := utils.GetAdminRouter().Group("styles")
	{
		rg.GET("/", renderList)
		rg.GET("/posts", renderForm)
		rg.GET("/edit", renderEditForm)
	}
}
