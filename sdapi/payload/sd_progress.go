package payload

type SdProgressState struct {
	Skipped       bool   `json:"skipped"`
	Interrupted   bool   `json:"interrupted"`
	Job           string `json:"job"`
	JobCount      int    `json:"job_count"`
	JobTimestamp  string `json:"job_timestamp"`
	JobNo         int    `json:"job_no"`
	SamplingStep  int    `json:"sampling_step"`
	SamplingSteps int    `json:"sampling_steps"`
}

type SdProgress struct {
	Progress     float32         `json:"progress"`
	EtaRelative  float32         `json:"eta_relative"`
	State        SdProgressState `json:"state"`
	CurrentImage string          `json:"current_image"`
	TextInfo     string          `json:"textinfo"`
}
