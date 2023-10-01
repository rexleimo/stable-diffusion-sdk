package category

import (
	"net/http"
	"stable-diffusion-sdk/admin/utils"

	"github.com/gin-gonic/gin"
)

func Init() {
	group := utils.GetAdminRouter().Group("category")
	group.GET("", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "category/index.html", gin.H{
			"title": "分类管理",
		})
	})
}
