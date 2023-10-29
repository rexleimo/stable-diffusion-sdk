package img2img

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"stable-diffusion-sdk/core/httpserver"
	"stable-diffusion-sdk/sdapi/handle"
	"stable-diffusion-sdk/sdapi/payload"

	"github.com/gin-gonic/gin"
)

func Init() {
	rg := httpserver.GetInstance().Group("sd")
	{
		rg.GET("img2img", func(ctx *gin.Context) {
			data, _ := ioutil.ReadFile("./public/sd_block/20231014/169721286640.png")
			base_64 := base64.StdEncoding.EncodeToString(data)

			json := payload.SDImageParams{
				OverrideSettings: payload.OverrideSettings{
					SdModelCheckpoint: "REX_majicMIX_realistic_v7.safetensors [7c819b6d13]",
				},
				Seed:                  -1,
				Width:                 632,
				Height:                1024,
				CfgScale:              7,
				Steps:                 20,
				Eta:                   0,
				BatchSize:             1,
				InitImages:            []string{base_64},
				Mask:                  "",
				SamplerName:           "DPM++ 2M Karras",
				DenoisingStrength:     0.75,
				InpaintFullResPadding: 32,
				AlwaysonScripts:       nil,
			}

			resp, _ := handle.Img2Imgapi(json)

			ctx.String(http.StatusOK, resp[0])
		})
	}
}
