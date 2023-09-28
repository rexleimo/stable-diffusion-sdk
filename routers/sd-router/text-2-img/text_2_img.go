package text2img

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"stable-diffusion-sdk/core/httpserver"
	"stable-diffusion-sdk/sdapi/handle"
	"stable-diffusion-sdk/sdapi/payload"
	"time"

	"github.com/gin-gonic/gin"
)

func Init() {
	group := httpserver.GetInstance().Group("sd")
	group.GET("/text2img", func(ctx *gin.Context) {
		json := &payload.SDParams{
			Seed:         -1,
			Width:        512,
			Height:       512,
			CfgScale:     7,
			Steps:        30,
			Eta:          0,
			SamplerIndex: "Euler",
			BatchSize:    1,
		}
		err := ctx.ShouldBindJSON(&json)

		if err != nil {
			ctx.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		fmt.Println(json)
		s, _ := handle.Text2ImgApi(*json)

		timestampFunc := func() int64 {
			return time.Now().UnixMicro()
		}

		image := make([]string, 0, 10)
		for _, v := range s {
			timestamp := timestampFunc()
			path := fmt.Sprintf("public/sd_block/%s/%d.png", time.Now().Format("20060102"), timestamp)
			image = append(image, path)

			go func(bStr string, p string) {
				b, _ := base64.StdEncoding.DecodeString(bStr)
				os.MkdirAll(filepath.Dir(p), 0755)
				err := os.WriteFile(p, b, 0644)
				if err != nil {
					fmt.Println(err)
				}
			}(v, path)

		}

		ctx.JSON(200, gin.H{
			"images": image,
		})
	})
}
