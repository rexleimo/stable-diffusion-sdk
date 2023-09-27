package payload

type OverrideSettings struct {
	SdModelCheckpoint string `json:"sd_model_checkpoint"` //模型名称
}

type SDParams struct {
	SamplerName       string           `json:"sampler_name"`    //采样器算法,不同的采样器效果也不同。
	Prompt            string           `json:"prompt"`          //生成图像的提示文本,描述你想要生成什么图像。
	NegativePrompt    string           `json:"negative_prompt"` //排除提示文本,描述你不想在图像中出现的内容。
	Seed              int64            `json:"seed"`            //随机种子数,用于控制生成结果。
	Width             int32            `json:"width"`
	Height            int32            `json:"height"`
	CfgScale          int32            `json:"cfg_scale"`          //调整生成器的参数,通常取值在7到12之间。
	Steps             int32            `json:"steps"`              // 生成步骤数,一般50就可以了。
	RestoreFaces      bool             `json:"restore_faces"`      //是否保留人脸。
	Tiling            bool             `json:"tiling"`             // 是否使用tiles。
	Eta               int32            `json:"eta"`                //调整生成器的速度和质量。
	DenoisingStrength int32            `json:"denoising_strength"` //降噪强度。
	SamplerIndex      string           `json:"sampler_index"`      //采样器的索引,比如"Euler"。
	BatchSize         int32            `json:"batch_size"`         //出图个数
	OverrideSettings  OverrideSettings `json:"override_settings"`  // sd模型
}

type SDResponse struct {
	Images []string `json:"images"`
	Info   string   `json:"info"`
}
