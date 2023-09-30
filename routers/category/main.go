package category

import (
	"stable-diffusion-sdk/core/httpserver"

	"github.com/gin-gonic/gin"
)

func Init() {
	group := httpserver.GetInstance().Group("category")
	group.GET("/", func(ctx *gin.Context) {})
	group.GET("/:id", func(ctx *gin.Context) {})

	group.POST("", create)
	group.POST("/:id/values", addValue)

	group.PUT("/", func(ctx *gin.Context) {})
	group.DELETE("/", func(ctx *gin.Context) {})
}
