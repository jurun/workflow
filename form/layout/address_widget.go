package layout

import (
	"fmt"
	"github.com/jurun/workflow/utils"
)

type AddressWidget struct {
	BaseWidget
	Value      *AddressValue `json:"value"`
	NeedDetail bool          `json:"need_detail"`
	Linkage    Linkage       `json:"linkage"`
}

func (this *AddressWidget) GetValue() interface{} {
	if this.Value == nil {
		this.Value = AddressValue{}.Init()
	}
	return this.Value
}

func (this *AddressWidget) SetValue(val interface{}) error {
	v := utils.NewConvert(val)
	var value AddressValue

	if err := v.Bind(&value); err != nil {
		return fmt.Errorf("字段: %s 的值不是有效地址格式, Err: %s", this.Field.GetLabel(), err.Error())
	}

	this.Value = &value
	return nil
}

func (this *AddressWidget) Diff(widget Widget) (Diff, bool) {
	if this.Value == nil && widget.GetValue() == nil {
		return Diff{}, false
	}
	diffValue, _ := widget.GetValue().(*AddressValue)

	if (this.Value.Province == diffValue.Province) && (this.Value.City == diffValue.City) && (this.Value.District == diffValue.District) && (this.Value.Detail == diffValue.Detail) {
		return Diff{}, false
	}

	records := make([]DiffItem, 0)
	records = append(records, DiffItem{
		Name:      this.Field.GetLabel(),
		FieldName: this.Field.GetName(),
		Type:      "change",
		Original:  this.Value.String(),
		Last:      diffValue.String(),
	})

	return records, true
}

func (this *AddressWidget) String() string {
	if this.Value == nil {
		return ""
	}

	return fmt.Sprintf("%s%s%s%s", this.Value.Province, this.Value.City, this.Value.District, this.Value.Detail)
}
