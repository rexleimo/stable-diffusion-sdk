package category

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"
	"time"

	"github.com/gin-gonic/gin"
)

func create(ctx *gin.Context) {

	var payload models.Categories
	if err := ctx.Bind(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	collection := mongodb.GetInstance().Collection(payload.TableName())
	payload.CreateAt = time.Now()
	payload.UpdateAt = time.Now()

	ior, err := collection.InsertOne(context.Background(), payload)
	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "created", "id": ior.InsertedID})
}
