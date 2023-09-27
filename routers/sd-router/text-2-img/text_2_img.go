package text2img

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	httpserver "stable-diffusion-sdk/core/http-server"
	"stable-diffusion-sdk/sdapi/handle"
	"time"

	"github.com/gin-gonic/gin"
)

func Init() {
	group := httpserver.GetInstance().Group("sd")
	group.GET("/text2img", func(ctx *gin.Context) {
		s, _ := handle.Text2ImgApi()
		// 便利s，将base64字符串写入到public/日前/时间戳.png

		image := make([]string, 0, 10)

		for _, v := range s {
			b, _ := base64.StdEncoding.DecodeString(v)
			path := fmt.Sprintf("public/%s/%d.png", time.Now().Format("20060102"), time.Now().Unix())
			os.MkdirAll(filepath.Dir(path), 0755)
			image = append(image, path)
			err := ioutil.WriteFile(path, b, 0644)
			fmt.Println(err)
		}

		ctx.JSON(200, gin.H{
			"images": image,
		})
	})
}
