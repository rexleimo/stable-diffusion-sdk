package styles

import (
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/sdapi/handle"

	"github.com/gin-gonic/gin"
)

func renderForm(ctx *gin.Context) {
	modles, _ := handle.GetSdModels()
	ctx.HTML(200, "styles/posts.html", gin.H{
		"models": modles,
	})
}

func renderEditForm(ctx *gin.Context) {
	idStr := ctx.Param("id")
	style, _ := handles.GetStyleOneById(idStr)
	modles, _ := handle.GetSdModels()
	ctx.HTML(200, "styles/edit.html", gin.H{
		"style":  style,
		"models": modles,
	})
}
