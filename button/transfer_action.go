package button

import (
	"github.com/jurun/workflow/organize"
)

type TransferAction struct {
	Label   string                    `json:"label"`
	Callers *organize.ProcessOrganize `json:"callers"`
}

func (this *TransferAction) GetLabel() string {
	return this.Label
}
