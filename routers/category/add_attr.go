package category

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"github.com/gin-gonic/gin"
)

func addAttr(ctx *gin.Context) {
	var json models.Attr
	err := ctx.ShouldBindJSON(&json)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c := mongodb.GetInstance().Collection(json.TableName())

	_, err2 := c.InsertOne(context.Background(), json)

	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err2.Error()})
		return
	}

	ctx.JSON(200, gin.H{"success": true})
}
