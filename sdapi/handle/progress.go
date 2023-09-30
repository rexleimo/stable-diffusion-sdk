package handle

import (
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/http"
)

func Progress() (*payload.SdProgress, error) {
	resp, err := http.GetSDServer().SetResult(&payload.SdProgress{}).Get("sdapi/v1/progress")
	if err != nil {
		return nil, err
	}

	apiResp := resp.Result().(*payload.SdProgress)

	return apiResp, nil
}
