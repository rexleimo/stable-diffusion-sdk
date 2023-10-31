package handle

import (
	"fmt"
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/http"
)

func Img2Imgapi(params payload.SDImageParams) ([]string, error) {
	resp, err := http.GetSDServer().SetResult(&payload.SDResponse{}).SetHeader("Content-Type", "application/json").SetBody(params).Post("sdapi/v1/img2img")
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())
	apiResp := resp.Result().(*payload.SDResponse)

	return apiResp.Images, nil
}
