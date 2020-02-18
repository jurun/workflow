package layout

import (
	"fmt"
	"github.com/cstockton/go-conv"
	"time"
)

type DateRangeWidget struct {
	BaseWidget
	Value       []int64 `json:"value"`
	Placeholder string  `json:"placeholder"`
	SelectType  string  `json:"select_type"`
	Format      string  `json:"format"`
	Linkage     Linkage `json:"linkage"`
	Formula     Formula `json:"formula"`
}

func (this *DateRangeWidget) GetValue() interface{} {
	if this.Value == nil {
		this.Value = make([]int64, 0)
	}

	return this.Value
}

func (this *DateRangeWidget) SetValue(val interface{}) error {
	if val == nil {
		return nil
	}
	vals, flag := val.([]interface{})
	if !flag {
		return fmt.Errorf("字段: %s 的值不是有效 Slice 格式", this.Field.GetLabel())
	}

	if len(vals) == 0 {
		return nil
	}

	this.Value = make([]int64, 0)
	for k, v := range vals {
		number, err := conv.Int64(v)
		if err != nil {
			return fmt.Errorf("字段: %s 的第 %d 个值不是有效时间戳格式", this.Field.GetLabel(), k)
		}

		this.Value = append(this.Value, number)
	}

	return nil
}

func (this *DateRangeWidget) Diff(widget Widget) (Diff, bool) {
	if this.Value == nil && widget.GetValue() == nil {
		return Diff{}, false
	}

	records := make([]DiffItem, 0)
	diff := widget.(*DateRangeWidget)
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

	if len(this.Value) == 0 {
		return nil, false
	}
	if (this.Value[0] != diff.Value[0]) || (this.Value[1] != diff.Value[1]) {
		records = append(records, DiffItem{
			Name:      this.Field.GetLabel(),
			FieldName: this.Field.GetName(),
			Type:      "change",
			Original:  this.value2string(),
			Last:      diff.value2string(),
		})
		return records, true
	}

	return records, false
}

func (this *DateRangeWidget) value2string() string {
	if this.Value == nil || len(this.Value) == 0 {
		return ""
	}

	start := time.Unix(0, this.Value[0]*int64(time.Millisecond))
	end := time.Unix(0, this.Value[1]*int64(time.Millisecond))
	str := fmt.Sprintf("%s - %s", start.Format("2006-02-01 15:04"), end.Format("2006-02-01 15:04"))
	return str
}

func (this *DateRangeWidget) String() string {
	return this.value2string()
}
