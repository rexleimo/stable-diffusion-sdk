package category

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func query(ctx *gin.Context) {

	pid := ctx.Query("pid")

	var list models.Categories
	c, err := mongodb.GetInstance().Collection(list.TableName()).Find(context.Background(), bson.D{
		{Key: "pid", Value: pid},
	})
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var data []models.Categories
	err = c.All(context.Background(), &data)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": data})
}
