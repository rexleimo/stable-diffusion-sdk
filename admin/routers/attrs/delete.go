package attrs

import (
	"stable-diffusion-sdk/handles"

	"github.com/gin-gonic/gin"
)

func delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	//	delete handle
	err := handles.DeleteAttrById(idStr)
	if err != nil {
		// json {error:''}
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": "delete attr success"})
}
