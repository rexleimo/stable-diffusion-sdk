package handle

import (
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/http"
)

func GetSdModels() ([]payload.SdModelsResponse, error) {
	r, err := http.GetSDServer().SetResult([]payload.SdModelsResponse{}).SetHeader("Content-Type", "application/json").Get("sdapi/v1/sd-models")
	if err != nil {
		return nil, err
	}
	i := r.Result().([]payload.SdModelsResponse)
	return i, nil
}
