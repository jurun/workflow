package button

type PauseAction struct {
	Label string `json:"label"`
}

func (this *PauseAction) GetLabel() string {
	return this.Label
}
