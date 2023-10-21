package handle

import (
	"encoding/json"
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/http"
)

func GetSdModels() ([]payload.SdModelsResponse, error) {
	r, err := http.GetSDServer().SetHeader("Content-Type", "application/json").Get("sdapi/v1/sd-models")
	if err != nil {
		return nil, err
	}
	i := r.String()
	var models []payload.SdModelsResponse
	err = json.Unmarshal([]byte(i), &models)
	if err != nil {
		return nil, err
	}
	return models, nil
}
