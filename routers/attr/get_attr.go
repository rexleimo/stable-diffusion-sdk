package attrs

import (
	"context"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func getAttrs(ctx *gin.Context) {
	cid := ctx.Query("cid")
	// cid nil error
	if cid == "" {
		ctx.JSON(400, gin.H{
			"error": "cid is nil",
		})
		return
	}

	var table models.Attr
	cur, err := mongodb.GetInstance().Collection(table.TableName()).Find(context.Background(), bson.D{
		{Key: "pid", Value: cid},
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	var data []models.Attr
	err2 := cur.All(context.Background(), &data)
	if err2 != nil {
		ctx.JSON(400, gin.H{
			"error": err2.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": data,
	})

}
