package handle

import (
	"fmt"
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/config"
	"stable-diffusion-sdk/utils/http"
)

func Text2ImgApi() ([]string, error) {
	params := &payload.SDParams{
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
	resp, err := http.GetInstance().R().SetResult(&payload.SDResponse{}).SetHeader("Content-Type", "application/json").SetBody(params).Post(fmt.Sprintf("%ssdapi/v1/txt2img", config.GetConfig().SDServer.Host))

	if err != nil {
		return nil, err
	}

	apiResp := resp.Result().(*payload.SDResponse)

	return apiResp.Images, nil
}
