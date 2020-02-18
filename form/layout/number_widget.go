package layout

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"github.com/jurun/workflow/utils"
)

type NumberWidget struct {
	BaseWidget
	Value          interface{} `json:"value"`
	Placholder     string      `json:"placholder"`
	Prefix         string      `json:"prefix"`
	Suffix         string      `json:"suffix"`
	AllowedDecimal bool        `json:"allowed_decimal"`
	MaxNumber      float64     `json:"max_number"`
	MinNumber      float64     `json:"min_number"`
	Linkage        Linkage     `json:"linkage"`
	Formula        Formula     `json:"formula"`
}

func (this *NumberWidget) GetValue() interface{} {
	//if this.Value == nil {
	//	return 0
	//}
	return this.Value
}

func (this *NumberWidget) SetValue(val interface{}) error {
	if val == nil {
		return nil
	}

	//number := utils.NewConvert(val).Float64(0)

	number, err := conv.Float64(val)
	if err != nil {
		return fmt.Errorf("字段: %s 的值不是有效的数字格式", this.Field.GetLabel())
	}

	this.Value = number
	return nil
}

func (this *NumberWidget) Diff(widget Widget) (Diff, bool) {
	if this.Value == nil && widget.GetValue() == nil {
		return Diff{}, false
	}

	//diff := widget.(*NumberWidget)

	origin := utils.NewConvert(this.Value).Float64(0)
	diff := utils.NewConvert(widget.GetValue()).Float64(0)
	if origin != diff {
		records := make([]DiffItem, 0)
		rc := DiffItem{
			Name:      this.Field.GetLabel(),
			FieldName: this.Field.GetName(),
			Type:      "change",
		}
		if this.Value == nil {
			rc.Original = ""
		} else {
			rc.Original = origin
		}
		if widget.GetValue() == nil {
			rc.Last = ""
		} else {
			rc.Last = diff
		}
		records = append(records, rc)
		return records, true
	}
	return Diff{}, false
}

func (this *NumberWidget) String() string {
	if this.Value == nil {
		return ""
	}

	return utils.NewConvert(this.Value).String()
}
