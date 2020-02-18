package layout

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"strings"
)

type CheckBoxWidget struct {
	Choice
	Value  []interface{} `json:"value"`
	Layout string        `json:"layout"`
}

func (this *CheckBoxWidget) GetValue() interface{} {
	if this.Value == nil {
		this.Value = []interface{}{}
	}

	return this.Value
}

func (this *CheckBoxWidget) SetValue(val interface{}) error {
	vals, flag := val.([]interface{})
	if !flag {
		return fmt.Errorf("字段: %s 的值不是有效的 Slice 格式", this.Field.GetLabel())
	}

	this.Value = make([]interface{}, 0)

	for _, v := range vals {
		number, err := conv.Float64(v)
		if err != nil {
			this.Value = append(this.Value, v)
		} else {
			this.Value = append(this.Value, number)
		}
	}

	return nil
}

func (this *CheckBoxWidget) Diff(widget Widget) (Diff, bool) {
	if this.Value == nil && widget.GetValue() == nil {
		return Diff{}, false
	}

	records := make([]DiffItem, 0)
	diff := widget.(*CheckBoxWidget)
	if len(this.Value) != len(diff.Value) {
		records = append(records, DiffItem{
			Name:      this.Field.GetLabel(),
			FieldName: this.Field.GetName(),
			Type:      "change",
			Original:  this.value2string(),
			Last:      diff.value2string(),
		})
		return records, true
	}

	for i := 0; i < len(this.Value); i++ {
		if this.Value[i] != diff.Value[i] {
			records = append(records, DiffItem{
				Name:      this.Field.GetLabel(),
				FieldName: this.Field.GetName(),
				Type:      "change",
				Original:  this.value2string(),
				Last:      diff.value2string(),
			})
			return records, true
		}
	}

	return records, false
}

func (this *CheckBoxWidget) value2string() string {
	if this.Value == nil || len(this.Value) == 0 {
		return ""
	}

	str := ""
	for _, v := range this.Value {
		str += fmt.Sprintf("%v,", v)
	}
	str = strings.TrimRight(str, ",")
	return str
}

func (this *CheckBoxWidget) String() string {
	return this.value2string()
}
