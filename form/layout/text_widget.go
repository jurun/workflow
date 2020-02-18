package layout

type TextWidget struct {
	BaseWidget
	Value       string  `json:"value"`
	Placeholder string  `json:"placeholder"`
	Prefix      string  `json:"prefix"`
	Suffix      string  `json:"suffix"`
	NoRepeat    bool    `json:"no_repeat"`
	Format      string  `json:"format"`
	Linkage     Linkage `json:"linkage"`
	Formula     Formula `json:"formula"`
}

func (this *TextWidget) GetValue() interface{} {
	if this.Value == "" {
		return nil
	}
	return this.Value
}

func (this *TextWidget) SetValue(val interface{}) error {
	str, flag := val.(string)
	if !flag {
		//return fmt.Errorf("字段: %s 的值不是有效 String 类型", this.Field.GetLabel())
	}

	this.Value = str
	return nil
}

func (this *TextWidget) Diff(widget Widget) (Diff, bool) {
	if this.Value == "" && widget.String() == "" {
		return nil, false
	}

	if this.Value != widget.(*TextWidget).Value {
		records := make([]DiffItem, 0)
		records = append(records, DiffItem{
			Name:      this.Field.GetLabel(),
			FieldName: this.Field.GetName(),
			Type:      "change",
			Original:  this.Value,
			Last:      widget.(*TextWidget).Value,
		})
		return records, true
	}
	return Diff{}, false
}

func (this *TextWidget) String() string {
	return this.Value
}
