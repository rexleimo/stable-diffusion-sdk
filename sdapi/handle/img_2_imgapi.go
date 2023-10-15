package handle

import (
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/http"

	"github.com/go-resty/resty/v2"
)

func Img2Imgapi(params payload.SDImageParams) *resty.Response {
	resp, _ := http.GetSDServer().SetResult(&payload.SDResponse{}).SetHeader("Content-Type", "application/json").SetBody(params).Post("sdapi/v1/img2img")
	return resp
}
