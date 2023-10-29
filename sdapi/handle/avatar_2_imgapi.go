package handle

import (
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"stable-diffusion-sdk/handles"
	"stable-diffusion-sdk/models"
	"stable-diffusion-sdk/sdapi/payload"
	"time"

	_ "golang.org/x/image/webp"
)

func GetImageSize(imagePath string) ([]int, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Println("1", err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("2", err)
	}
	return []int{img.Bounds().Dx(), img.Bounds().Dy()}, nil
}

func AvatarProgress(task models.Task) ([]string, error) {
	fmt.Println("render avatar")
	fmt.Println(task.QrcodePath)
	data, _ := ioutil.ReadFile(task.QrcodePath)
	size, err := GetImageSize(task.QrcodePath)
	fmt.Println(err)
	inputImage := base64.StdEncoding.EncodeToString(data)

	styleEntity, err := handles.GetStyleOneById(task.CID)
	if err != nil {
		return nil, err
	}

	json := payload.SDImageParams{
		Prompt:         styleEntity.Pormpt,
		NegativePrompt: styleEntity.NegativePrompt,
		OverrideSettings: payload.OverrideSettings{
			SdModelCheckpoint: styleEntity.Checkpoint,
		},
		Seed:                  -1,
		Width:                 int32(size[0]),
		Height:                int32(size[1]),
		CfgScale:              int32(styleEntity.CfgScale),
		Steps:                 int32(styleEntity.Steps),
		Eta:                   0,
		BatchSize:             1,
		SamplerIndex:          styleEntity.SamplerIndex,
		InitImages:            []string{inputImage},
		Mask:                  "",
		DenoisingStrength:     0.45,
		InpaintFullResPadding: 32,
		AlwaysonScripts: &payload.AlwaysonScripts{
			Controlnet: payload.Controlnet{
				Args: []payload.ControlnetArg{
					{
						Enable:        true,
						Module:        "canny",
						Model:         "control_v11p_sd15_canny [d14c016b]",
						ResizeMode:    1,
						ProcessorRes:  size[0],
						Weight:        1.4,
						GuidanceStart: 0,
						GuidanceEnd:   1,
					},
				},
			},
		},
	}

	s, _ := Img2Imgapi(json)

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
