package category

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context) {
	var json models.Categories
	if err := ctx.BindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	collection := mongodb.GetInstance().Collection(json.TableName())

	ior, err := collection.InsertOne(context.Background(), json)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "created", "id": ior.InsertedID})

}
