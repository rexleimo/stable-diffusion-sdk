package handle

import (
	"fmt"
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/http"
)

func Text2ImgApi() {
	payload := &payload.SDParams{
		Prompt:         "car",
		NegativePrompt: "",
		OverrideSettings: payload.OverrideSettings{
			SdModelCheckpoint: "unstableinkdream_v80Photo.safetensors [bd67d37dac]",
		},
		Seed:         -1,
		Width:        512,
		Height:       512,
		CfgScale:     7,
		Steps:        30,
		Eta:          0,
		SamplerIndex: "Euler",
		BatchSize:    1,
	}
	fmt.Println(payload)
	resp, _ := http.GetInstance().R().SetHeader("Content-Type", "application/json").SetBody(payload).Post("sdapi/v1/txt2img")
	fmt.Println(resp.String())
}
