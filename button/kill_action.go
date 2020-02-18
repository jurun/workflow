package button

type KillAction struct {
	Label string `json:"label"`
}

func (this *KillAction) GetLabel() string {
	return this.Label
}
