package qrcode

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"stable-diffusion-sdk/sdapi/payload"
	"time"

	"github.com/gin-gonic/gin"
)

func progressQrcode(ctx *gin.Context) {
	// data, _ := ioutil.ReadFile("./public/sd_block/test.jpg")
	// base64img := base64.StdEncoding.EncodeToString(data)
	var row payload.SDQrcodeParams
	err := ctx.ShouldBindJSON(&row)

	if err != nil {
		// response json {error:""}
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	path := fmt.Sprintf("public/sd_block/%s/%d.png", time.Now().Format("20060102"), time.Now().UnixMicro())
	data, _ := base64.StdEncoding.DecodeString(row.Qrcode)
	os.MkdirAll(filepath.Dir(path), 0755)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"success": "ok"})

}
