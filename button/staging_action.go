package button

type StagingAction struct {
	Label string `json:"label"`
}

func (this *StagingAction) GetLabel() string {
	return this.Label
}
