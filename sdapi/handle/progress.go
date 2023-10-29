package handle

import (
	"fmt"
	"stable-diffusion-sdk/models"
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

func ProcessTask(task models.Task) ([]string, error) {
	fmt.Println("TaskType:", task)
	if task.Type == 0 {
		return Text2ImgProcess(task)
	} else if task.Type == 1 {
		return nil, nil
	} else if task.Type == 4 {
		return AvatarProgress(task)
	} else {
		return QrcodeProcess(task)
	}
}
