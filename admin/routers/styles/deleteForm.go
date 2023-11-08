package styles

import (
	"stable-diffusion-sdk/handles"

	"github.com/gin-gonic/gin"
)

func deleteForm(ctx *gin.Context) {
	idStr := ctx.Param("id")
	_, err := handles.DeleteStyle(idStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": "ok"})
}
