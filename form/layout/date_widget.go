package layout

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"time"
)

type DateWidget struct {
	BaseWidget
	Value       int64   `json:"value"`
	Placeholder string  `json:"placeholder"`
	SelectType  string  `json:"select_type"`
	Format      string  `json:"format"`
	Linkage     Linkage `json:"linkage"`
	Formula     Formula `json:"formula"`
}

func (this *DateWidget) GetValue() interface{} {
	return this.Value
}

func (this *DateWidget) SetValue(val interface{}) error {
	if val == nil {
		return nil
	}

	number, err := conv.Int64(val)
	if err != nil {
		return fmt.Errorf("字段: %s 的值不是有效的时间戳", this.Field.GetLabel())
	}

	this.Value = number
	return nil
}

func (this *DateWidget) Diff(widget Widget) (Diff, bool) {
	records := make([]DiffItem, 0)
	if this.Value != widget.(*DateWidget).Value {
		re := DiffItem{
			Name:      this.Field.GetLabel(),
			FieldName: this.Field.GetName(),
			Type:      "change",
		}
		if this.Value == 0 {
			re.Original = ""
		} else {
			re.Original = time.Unix(0, this.Value*int64(time.Millisecond)).Format("2006-02-01 15:04")
		}

		if widget.(*DateWidget).Value == 0 {
			re.Last = ""
		} else {
			re.Last = time.Unix(0, widget.(*DateWidget).Value*int64(time.Millisecond)).Format("2006-02-01 15:04")
		}

		records = append(records, re)
		//records = append(records, form_common.DiffItem{
		//	Name:     this.Field.GetLabel(),
		//	Type:     "change",
		//	Original: fmt.Sprintf("%d", this.Value),
		//	Last:     fmt.Sprintf("%d", widget.(*DateWidget).Value),
		//})
		return records, true
	}

	return records, false
}

func (this *DateWidget) String() string {
	if this.Value == 0 {
		return ""
	}

	return time.Unix(0, this.Value*int64(time.Millisecond)).Format("2006-02-01 15:04")
}
