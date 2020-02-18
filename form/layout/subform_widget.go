package layout

import (
	"fmt"
	"github.com/jurun/workflow/utils"
)

type SubFormWidget struct {
	BaseWidget
	Items []interface{}            `json:"items"`
	Value []map[string]interface{} `json:"value"`
}

func (this *SubFormWidget) GetValue() interface{} {
	if this.Value == nil {
		this.Value = make([]map[string]interface{}, 0)
	}

	return this.Value
}

func (this *SubFormWidget) SetValue(val interface{}) error {
	this.Value = make([]map[string]interface{}, 0)
	values := make([]map[string]interface{}, 0)

	err := utils.NewConvert(val).Bind(&values)
	if err != nil {
		return fmt.Errorf("字段: %s 的值格式不正确", this.Field.GetLabel())
	}

	//subFields := this.Field.GetSubFields()

	for _, v := range values {
		subData := make(map[string]interface{})
		for fieldName, fieldValue := range v {
			field, found := this.findField(fieldName)
			if !found {
				continue
				//return fmt.Errorf("字段: %s 中未找到子字段: %s", this.Field.GetLabel(), fieldName)
			}

			wgt, err := field.GetWidget()
			if err != nil {
				return fmt.Errorf("字段: %s 的结构错误", field.GetLabel())
			}

			if err = wgt.SetValue(fieldValue); err != nil {
				return err
			}

			subData[fieldName] = wgt.GetValue()
		}

		this.Value = append(this.Value, subData)
	}

	return nil
}

func (this *SubFormWidget) Diff(widget Widget) (Diff, bool) {
	if this.Value == nil && widget.GetValue() == nil {
		return nil, false
	}

	diff := widget.(*SubFormWidget)
	records := make([]DiffItem, 0)

	if len(this.Value) != len(diff.Value) {
		records = append(records, DiffItem{
			Name:      this.Field.GetLabel(),
			FieldName: this.Field.GetName(),
			Type:      "change",
			Original:  "",
			Last:      "",
		})
		return records, true
	}

	// @todo 调子 Field 去比较
	//for i := 0; i < len(this.Value); i++ {
	//	for k, v := range this.Value[i] {
	//		if v != diff.Value[i][k] {
	//			records = append(records, form_common.DiffItem{
	//				Name:     this.Field.GetLabel(),
	//				Type:     "change",
	//				Original: "",
	//				Last:     "",
	//			})
	//			return records, true
	//		}
	//	}
	//}

	//for i := 0; i < len(this.Value); i++ {
	//	values := this.Value[i]
	//	diffValues := diff.Value[i]
	//
	//	for key, val := range values {
	//		diffVal, flag := diffValues[key]
	//		if !flag {
	//			records = append(records, form_common.DiffItem{
	//				Name: this.Field.GetLabel(),
	//				Type: "change",
	//			})
	//			return records, true
	//		}
	//
	//		if val != diffVal {
	//			records = append(records, form_common.DiffItem{
	//				Name: this.Field.GetLabel(),
	//				Type: "change",
	//			})
	//			return records, true
	//		}
	//	}
	//}

	return Diff{}, false
}

func (this *SubFormWidget) findField(name string) (*Field, bool) {
	for _, f := range this.Field.GetSubFields() {
		if f.GetName() == name {
			return f, true
		}
	}
	return nil, false
}

//func (this *SubFormWidget) value2strings() []string  {
//	strs:=make([]string,0)
//	if this.Value == nil || len(this.Value) == 0 {
//		return strs
//	}
//
//	for i := 0; i < len(this.Value); i++ {
//		strs = append(strs,this.value2string(i))
//	}
//	return strs
//}

func (this *SubFormWidget) value2string(index int, subFieldName string) string {
	if this.Value == nil || len(this.Value) == 0 {
		return ""
	}

	if index > len(this.Value)-1 {
		return ""
	}

	value, flag := this.Value[index][subFieldName]
	if !flag {
		return ""
	}
	return fmt.Sprintf("%v", value)
}

func (this *SubFormWidget) String() string {
	if this.Value == nil || len(this.Value) == 0 {
		return ""
	}

	text := ""
	for _, v := range this.Value {
		for fieldName, val := range v {
			field, found := this.findField(fieldName)
			if !found {
				continue
			}

			wgt, err := field.GetWidget()
			if err != nil {
				continue
			}

			wgt.SetValue(val)
			text += wgt.String()
		}
	}

	return text
}
