package handle

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/sdapi/payload"
	"time"
)

func QrcodeProcess(task models.Task) ([]string, error) {
	fmt.Println("render qrcode")
	data, _ := ioutil.ReadFile(task.QrcodePath)
	inputImage := base64.StdEncoding.EncodeToString(data)

	styleEntity, err := handles.GetStyleOneById(task.CID)
	if err != nil {
		return nil, err
	}

	json := payload.SDParams{
		Prompt:         styleEntity.Pormpt,
		NegativePrompt: styleEntity.NegativePrompt,
		OverrideSettings: payload.OverrideSettings{
			SdModelCheckpoint: styleEntity.Checkpoint,
		},
		Seed:         -1,
		Width:        768,
		Height:       768,
		CfgScale:     int32(styleEntity.CfgScale),
		Steps:        int32(styleEntity.Steps),
		Eta:          0,
		BatchSize:    1,
		SamplerIndex: styleEntity.SamplerIndex,
		AlwaysonScripts: &payload.AlwaysonScripts{
			Controlnet: payload.Controlnet{
				Args: []payload.ControlnetArg{
					{
						Enable:        true,
						InputImage:    inputImage,
						Module:        "none",
						Model:         "control_v1p_sd15_qrcode_monster [a6e58995]",
						ResizeMode:    1,
						Weight:        1.6,
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
