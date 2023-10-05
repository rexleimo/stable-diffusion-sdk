package category

import (
	"context"
	"fmt"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/utils/mongodb"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func update(ctx *gin.Context) {

	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(400, gin.H{"error": "id is required"})
		return
	}

	id, _ := primitive.ObjectIDFromHex(idParam)
	var payload models.Categories
	if err := ctx.ShouldBind(&payload); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	collection := mongodb.GetInstance().Collection(payload.TableName())
	payload.CreateAt = time.Now()
	payload.UpdateAt = time.Now()

	fmt.Println(payload)

	_, err2 := collection.UpdateOne(context.Background(), bson.D{{Key: "_id", Value: id}}, bson.D{{Key: "$set", Value: payload}})
	if err2 != nil {
		fmt.Println(err2)
		ctx.JSON(401, gin.H{"error": err2.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "success"})
}
