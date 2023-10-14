package taskroute

import (
	"stable-diffusion-sdk/handles"

	"github.com/gin-gonic/gin"
)

func delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userIdStr := ctx.GetString("user_id")
	err := handles.DeleteTask(idStr, userIdStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error(), "error_code": 10000})
		return
	}
	ctx.JSON(200, gin.H{"success": true})
}
