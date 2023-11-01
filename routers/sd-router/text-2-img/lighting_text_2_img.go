package text2img

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

func lightingText2Img(ctx *gin.Context) {
	userID := ctx.GetString("user_id")
	userEntity, err := handles.FindUserById(userID)
	if err != nil {
		// response json error
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var payload payload.SDQrcodeParams
	err2 := ctx.ShouldBindJSON(&payload)
	if err2 != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	path := fmt.Sprintf("public/sd_block/qrcode/%s/%d.png", time.Now().Format("20060102"), time.Now().UnixMicro())
	data, _ := base64.StdEncoding.DecodeString(payload.Qrcode)
	os.MkdirAll(filepath.Dir(path), 0755)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	taskEntity := models.Task{
		Type:       5,
		CID:        payload.StyleID,
		UID:        userID,
		QrcodePath: path,
	}

	ior, err3 := handles.InsertTask(taskEntity)
	if err3 != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userEntity.Bonus -= 1
	handles.UpdateUser(userEntity)

	taskEntity.ID = ior.InsertedID.(primitive.ObjectID)
	go queue.InstallQueyue(taskEntity)

	ctx.JSON(200, gin.H{"ior": ior.InsertedID})

}
