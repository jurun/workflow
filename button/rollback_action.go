package button

type RollbackAction struct {
	Label string `json:"label"`
	Range []int  `json:"range"`
	Prev  bool   `json:"prev"`
}

func (this *RollbackAction) GetLabel() string {
	return this.Label
}
