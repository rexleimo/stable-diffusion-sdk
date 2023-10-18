package handle

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/sdapi/payload"
	"stable-diffusion-sdk/utils/http"
	"time"
)

func Text2ImgApi(params payload.SDParams) ([]string, error) {

	resp, err := http.GetSDServer().SetResult(&payload.SDResponse{}).SetHeader("Content-Type", "application/json").SetBody(params).Post("sdapi/v1/txt2img")

	if err != nil {
		return nil, err
	}

	apiResp := resp.Result().(*payload.SDResponse)

	return apiResp.Images, nil
}

func Text2ImgProcess(task models.Task) ([]string, error) {
	json := task.SDRendererConfig
	categroy, _ := handles.GetCategoryById(task.CID)
	json.OverrideSettings.SdModelCheckpoint = categroy.Checkpoint
	json.CfgScale = categroy.CfgScale
	json.Steps = categroy.Steps
	json.SamplerName = categroy.SamplerIndex

	if categroy.IsSize {
		json.Width = categroy.Imgw
		json.Height = categroy.Imgh
	}

	json.NegativePrompt = fmt.Sprintf(`%s,%s`, json.NegativePrompt, categroy.NegativePrompt)
	json.Prompt = fmt.Sprintf(`%s,%s`, json.Prompt, categroy.Pormpt)

	// log the NegativePrompt and Prompt
	fmt.Println("NegativePrompt:", json.NegativePrompt)
	fmt.Println("Prompt:", json.Prompt)

	s, _ := Text2ImgApi(json)

	timestampFunc := func() string {
		return fmt.Sprintf("%d%d", time.Now().Unix(), rand.Intn(1000))
	}

	image := make([]string, 0, 10)
	for _, v := range s {
		timestamp := timestampFunc()
		path := fmt.Sprintf("public/sd_block/%s/%s.png", time.Now().Format("20060102"), timestamp)
		image = append(image, path)

		go func(bStr string, p string) {
			b, _ := base64.StdEncoding.DecodeString(bStr)
			os.MkdirAll(filepath.Dir(p), 0755)
			err := os.WriteFile(p, b, 0644)
			if err != nil {
				fmt.Println(err)
			}
		}(v, path)
	}
	fmt.Println(image)
	return image, nil
}
