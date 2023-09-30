package attrvalues

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getValues(ctx *gin.Context) {
	aId := ctx.Query("aid")
	if aId == "" {
		ctx.JSON(400, gin.H{"error": "aid is required"})
	}
	var table models.AttrValue
	cur, err := mongodb.GetInstance().Collection(table.TableName()).Find(context.Background(), bson.D{
		{Key: "pid", Value: aId},
	})

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var values []models.AttrValue
	err2 := cur.All(context.Background(), &values)
	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err2.Error()})
		return
	}
	// TODO: pagination
	ctx.JSON(200, gin.H{"data": values})
}
