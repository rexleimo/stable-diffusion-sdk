package handle

import (
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/http"
)

func Text2ImgApi(params payload.SDParams) ([]string, error) {
	// params := &payload.SDParams{
	// 	Prompt:         "car",
	// 	NegativePrompt: "",
	// 	OverrideSettings: payload.OverrideSettings{
	// 		SdModelCheckpoint: "unstableinkdream_v80Photo.safetensors [bd67d37dac]",
	// 	},
	// 	Seed:         -1,
	// 	Width:        512,
	// 	Height:       512,
	// 	CfgScale:     7,
	// 	Steps:        30,
	// 	Eta:          0,
	// 	SamplerIndex: "Euler",
	// 	BatchSize:    1,
	// }
	resp, err := http.GetSDServer().SetResult(&payload.SDResponse{}).SetHeader("Content-Type", "application/json").SetBody(params).Post("sdapi/v1/txt2img")

	if err != nil {
		return nil, err
	}

	apiResp := resp.Result().(*payload.SDResponse)

	return apiResp.Images, nil
}
