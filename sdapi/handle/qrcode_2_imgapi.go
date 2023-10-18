package handle

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/sdapi/payload"
	"time"
)

func QrcodeProcess(task models.Task) ([]string, error) {
	fmt.Println("render qrcode")
	data, _ := ioutil.ReadFile(task.QrcodePath)
	inputImage := base64.StdEncoding.EncodeToString(data)

	json := payload.SDParams{
		Prompt:         "interiortinyhouse interior couch, kitchen, wooden, stairs, table, stove, pan, mug, rug, ((masterpiece)), realistic, epic, details,<lora:ARWinteriortinyhouse:1>",
		NegativePrompt: "bad-picture-chill-75v,  badhandsv5-neg,  badhandv4,  By bad artist -neg,  easynegative,  ng_deepnegative_v1_75t,  verybadimagenegative_v1.1-6400, Watermark, Text, censored, deformed, bad anatomy, disfigured, poorly drawn face, mutated, extra limb, ugly, poorly drawn hands, missing limb, floating limbs, disconnected limbs, disconnected head, malformed hands, long neck, mutated hands and fingers, bad hands, missing fingers, cropped, worst quality, low quality, mutation, poorly drawn, huge calf, bad hands, fused hand, missing hand, disappearing arms, disappearing thigh, disappearing calf, disappearing legs, missing fingers, fused fingers, abnormal eye proportion, Abnormal hands, abnormal legs, abnormal feet,  abnormal fingers",
		OverrideSettings: payload.OverrideSettings{
			SdModelCheckpoint: "xxmix9realistic_v40.safetensors [18ed2b6c48]",
		},
		Seed:        -1,
		Width:       768,
		Height:      768,
		CfgScale:    7,
		Steps:       20,
		Eta:         0,
		BatchSize:   1,
		SamplerName: "Euler a",
		AlwaysonScripts: payload.AlwaysonScripts{
			Controlnet: payload.Controlnet{
				Args: []payload.ControlnetArg{
					{
						Enable:        true,
						InputImage:    inputImage,
						Module:        "none",
						Model:         "control_v1p_sd15_qrcode_monster [a6e58995]",
						ResizeMode:    1,
						Weight:        1.5,
						GuidanceStart: 0,
						GuidanceEnd:   1,
					},
				},
			},
		},
	}

	s, _ := Text2ImgApi(json)

	timestampFunc := func() string {
		return fmt.Sprintf("%d%d", time.Now().Unix(), rand.Intn(1000))
	}

	image := make([]string, 0, 10)

	path := fmt.Sprintf("public/sd_block/%s/%s.png", time.Now().Format("20060102"), timestampFunc())
	image = append(image, path)

	go func(bStr string, p string) {
		b, _ := base64.StdEncoding.DecodeString(bStr)
		os.MkdirAll(filepath.Dir(p), 0755)
		err := os.WriteFile(p, b, 0644)
		if err != nil {
			fmt.Println(err)
		}
	}(s[0], path)

	fmt.Println(image)
	return image, nil
}
