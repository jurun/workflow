package layout

import (
	"fmt"
	"github.com/cstockton/go-conv"
)

type SelectWidget struct {
	Choice
	Value       interface{} `json:"value"`
	Placeholder string      `json:"placeholder"`
	Search      bool        `json:"search"`
}

func (this *SelectWidget) GetValue() interface{} {
	return this.Value
}

func (this *SelectWidget) SetValue(val interface{}) error {
	number, err := conv.Float64(val)
	if err != nil {
		this.Value = val
	} else {
		this.Value = number
	}
	return nil
}

func (this *SelectWidget) Diff(widget Widget) (Diff, bool) {
	if this.Value == nil && widget.GetValue() == nil {
		return Diff{}, false
	}

	if this.String() == "" && widget.String() == "" {
		return nil, false
	}

	records := make([]DiffItem, 0)
	if this.Value != widget.GetValue() {
		re := DiffItem{
			Name:      this.Field.GetLabel(),
			FieldName: this.Field.GetName(),
			Type:      "change",
		}
		if this.Value == nil {
			re.Original = ""
		} else {
			re.Original = fmt.Sprintf("%v", this.Value)
		}

		if widget.GetValue() == nil {
			re.Last = ""
		} else {
			re.Last = fmt.Sprintf("%v", widget.GetValue())
		}

		records = append(records, re)
		return records, true
	}
	return Diff{}, false
}

func (this *SelectWidget) String() string {
	if this.Value == nil {
		return ""
	}
	return fmt.Sprintf("%v", this.Value)
}
