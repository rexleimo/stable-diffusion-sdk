package handle

import (
	"encoding/base64"
	"encoding/json"
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

	// params to json string
	paramsJson, _ := json.Marshal(params)
	// print the json
	fmt.Println(string(paramsJson))

	resp, err := http.GetSDServer().SetResult(&payload.SDResponse{}).SetHeader("Content-Type", "application/json").SetBody(params).Post("sdapi/v1/txt2img")

	if err != nil {
		return nil, err
	}
	fmt.Println(resp.String())
	apiResp := resp.Result().(*payload.SDResponse)

	return apiResp.Images, nil
}

func Text2ImgProcess(task models.Task) ([]string, error) {
	var payload payload.SDParams
	json := task.SDRendererConfig
	categroy, _ := handles.GetCategoryById(task.CID)
	payload.OverrideSettings.SdModelCheckpoint = categroy.Checkpoint
	payload.CfgScale = categroy.CfgScale
	payload.Steps = categroy.Steps
	payload.SamplerIndex = categroy.SamplerIndex

	if categroy.IsSize {
		payload.Width = categroy.Imgw
		payload.Height = categroy.Imgh
	}

	payload.NegativePrompt = fmt.Sprintf(`%s,%s`, json.NegativePrompt, categroy.NegativePrompt)
	payload.Prompt = fmt.Sprintf(`%s,%s`, json.Prompt, categroy.Pormpt)
	payload.BatchSize = 1
	payload.Seed = -1
	payload.DenoisingStrength = 7.0

	payload.AlwaysonScripts = nil

	s, _ := Text2ImgApi(payload)

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
