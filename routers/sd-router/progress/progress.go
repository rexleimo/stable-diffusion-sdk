package progress

import (
	"stable-diffusion-sdk/core/httpserver"
	"stable-diffusion-sdk/sdapi/handle"

	"github.com/gin-gonic/gin"
)

func Init() {
	group := httpserver.GetInstance().Group("sd")
	group.GET("/progress", func(ctx *gin.Context) {
		sp, err := handle.Progress()
		if err != nil {
			ctx.JSON(500, err)
			return
		}

		ctx.JSON(200, gin.H{
			"progress":      sp.Progress,
			"state":         sp.State,
			"eta_relative":  sp.EtaRelative,
			"current_image": sp.CurrentImage,
			"textinfo":      sp.TextInfo,
		})
	})
}
