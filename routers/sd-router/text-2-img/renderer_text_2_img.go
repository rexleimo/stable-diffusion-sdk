package text2img

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/sdapi/handle"
	"stable-diffusion-sdk/sdapi/payload"
	"time"

	"github.com/gin-gonic/gin"
)

func sendText2Img(ctx *gin.Context) {
	cid := ctx.Query("cid")

	if cid == "" {
		ctx.JSON(400, gin.H{"error": "cid is empty"})
		return
	}

	json := &payload.SDParams{
		Seed:     -1,
		Width:    1024,
		Height:   1024,
		CfgScale: 7,
		Steps:    30,
		Eta:      0,
		// SamplerIndex: "DPM++ 2M Karras",
		BatchSize: 1,
		// OverrideSettings: payload.OverrideSettings{
		// 	SdModelCheckpoint: "realisticVisionV51_v40VAE.safetensors [e9d3cedc4b]",
		// },
	}

	categroy, _ := handles.GetCategoryById(cid)
	json.OverrideSettings.SdModelCheckpoint = categroy.Checkpoint
	json.CfgScale = categroy.CfgScale
	json.Steps = categroy.Steps
	json.SamplerName = categroy.SamplerIndex

	if categroy.IsSize {
		json.Width = categroy.Imgw
		json.Height = categroy.Imgh
	}

	err := ctx.ShouldBindJSON(&json)

	json.NegativePrompt = fmt.Sprintf(`%s,%s`, json.NegativePrompt, categroy.NegativePrompt)
	json.Prompt = fmt.Sprintf(`%s,%s`, json.Prompt, categroy.Pormpt)
	fmt.Println(json)

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	s, _ := handle.Text2ImgApi(*json)

	timestampFunc := func() string {
		return fmt.Sprintf("%d%d", time.Now().Unix(), rand.Intn(1000))
	}

	image := make([]string, 0, 10)
	for _, v := range s {
		timestamp := timestampFunc()
		path := fmt.Sprintf("public/sd_block/%s/%s.png", time.Now().Format("20060102"), timestamp)
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
}
