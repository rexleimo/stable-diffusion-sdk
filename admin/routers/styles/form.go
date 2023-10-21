package styles

import (
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"

	"github.com/gin-gonic/gin"
)

func postForm(ctx *gin.Context) {
	var style models.Style
	err := ctx.ShouldBindJSON(&style)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// TODO: validate style
	ior, err2 := handles.InsertStyle(style)
	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err2.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": true, "insert_id": ior})
}

func editForm(ctx *gin.Context) {

	idStr := ctx.Param("id")

	var style models.Style
	err := ctx.ShouldBindJSON(&style)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// TODO: validate style
	ur, err2 := handles.UpdateStyleById(idStr, style)
	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err2.Error()})
		return
	}
	ctx.JSON(200, gin.H{"success": true, "update_id": ur.UpsertedID})
}
