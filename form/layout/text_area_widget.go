package layout

type TextAreaWidget struct {
	BaseWidget
	Value       string  `json:"value"`
	Placeholder string  `json:"placeholder"`
	Linkage     Linkage `json:"linkage"`
	Formula     Formula `json:"formula"`
}

func (this *TextAreaWidget) GetValue() interface{} {
	if this.Value == "" {
		return nil
	}

	return this.Value
}

func (this *TextAreaWidget) SetValue(val interface{}) error {
	str, flag := val.(string)
	if !flag {
		//return fmt.Errorf("字段: %s 的值不是有效 String 类型", this.Field.GetLabel())
	}

	this.Value = str
	return nil
}

func (this *TextAreaWidget) Diff(widget Widget) (Diff, bool) {
	if this.String() == "" && widget.String() == "" {
		return nil, false
	}

	if this.Value != widget.(*TextAreaWidget).Value {
		records := make([]DiffItem, 0)
		records = append(records, DiffItem{
			Name:      this.Field.GetLabel(),
			FieldName: this.Field.GetName(),
			Type:      "change",
			Original:  this.Value,
			Last:      widget.(*TextAreaWidget).Value,
		})
	}
	return Diff{}, false
}

func (this *TextAreaWidget) String() string {
	return this.Value
}
