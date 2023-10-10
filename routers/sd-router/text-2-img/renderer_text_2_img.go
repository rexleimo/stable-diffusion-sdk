package text2img

import (
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/queue"
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/mongodb"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func sendText2ImgTask(ctx *gin.Context) {
	cid := ctx.Query("cid")

	if cid == "" {
		ctx.JSON(400, gin.H{"error": "cid is empty"})
		return
	}

	json := payload.SDParams{
		Seed:      -1,
		Width:     1024,
		Height:    1024,
		CfgScale:  7,
		Steps:     30,
		Eta:       0,
		BatchSize: 1,
	}

	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userID, _ := ctx.Get("user_id")

	insertData := models.Task{
		UID:              userID.(string),
		CID:              cid,
		SDRendererConfig: json,
		CreateAt:         time.Now(),
		UpdateAt:         time.Now(),
	}

	c := mongodb.GetInstance().Collection(insertData.TableName())
	ior, err := c.InsertOne(ctx.Request.Context(), insertData)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	insertData.ID = ior.InsertedID.(primitive.ObjectID)
	go func(data models.Task) {
		queue.RendererTaskChan() <- data
	}(insertData)
	ctx.JSON(200, gin.H{"result": ior.InsertedID})
}
