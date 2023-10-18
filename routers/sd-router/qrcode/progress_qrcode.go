package qrcode

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/queue"
	"stable-diffusion-sdk/sdapi/payload"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func progressQrcode(ctx *gin.Context) {
	// data, _ := ioutil.ReadFile("./public/sd_block/test.jpg")
	// base64img := base64.StdEncoding.EncodeToString(data)

	userId := ctx.GetString("user_id")

	var row payload.SDQrcodeParams
	err := ctx.ShouldBindJSON(&row)

	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	path := fmt.Sprintf("public/sd_block/qrcode/%s/%d.png", time.Now().Format("20060102"), time.Now().UnixMicro())
	data, _ := base64.StdEncoding.DecodeString(row.Qrcode)
	os.MkdirAll(filepath.Dir(path), 0755)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	insertData := models.Task{
		Type:       3,
		UID:        userId,
		CID:        row.StyleID,
		QrcodePath: path,
	}

	ior, err := handles.InsertTask(insertData)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	insertData.ID = ior.InsertedID.(primitive.ObjectID)
	go func(data models.Task) {
		queue.RendererTaskChan() <- data
	}(insertData)

	ctx.JSON(200, gin.H{"success": "ok"})

}
