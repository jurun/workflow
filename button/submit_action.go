package button

type SubmitAction struct {
	Label string `json:"label"`
}

func (this *SubmitAction) GetLabel() string {
	return this.Label
}
