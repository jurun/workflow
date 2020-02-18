package button

type SubmitPrintAction struct {
	Label      string `json:"label"`
	TemplateId int    `json:"template_id"`
}

func (this *SubmitPrintAction) GetLabel() string {
	return this.Label
}
