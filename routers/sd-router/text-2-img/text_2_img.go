package text2img

import (
	httpserver "stable-diffusion-sdk/core/http-server"
	"stable-diffusion-sdk/sdapi/handle"

	"github.com/gin-gonic/gin"
)

func Init() {
	group := httpserver.GetInstance().Group("sd")
	group.GET("/text2img", func(ctx *gin.Context) {
		s, _ := handle.Text2ImgApi()
		ctx.JSON(200, gin.H{
			"images": s,
		})
	})
}
