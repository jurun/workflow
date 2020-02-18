package layout

type Tab struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Layout struct {
	Fields []*Field `json:"fields"`
	Tabs   []*Tab   `json:"tabs"`
}

func (Layout) Init() *Layout {
	return &Layout{
		Fields: []*Field{},
		Tabs:   []*Tab{},
	}
}

func (this *Layout) Mobile() {
	fields := this.AllFields(true)
	for _, f := range fields {
		f.SetAuth(FieldAuth{
			View:     true,
			Edit:     true,
			Required: true,
		})
	}
	this.Fields = fields
}

func (this *Layout) FindField(name string, mobile ...bool) (*Field, bool) {
	for _, f := range this.AllFields(mobile...) {
		if f.Name == name {
			return f, true
		}
	}
	return &Field{}, false
}

func (this *Layout) AllFields(mobile ...bool) []*Field {
	var flag bool
	if len(mobile) > 0 {
		flag = mobile[0]
	}

	fields := make([]*Field, 0)
	for _, f := range this.Fields {
		if f.Type == SEPARATOR {
			continue
		}

		//if flag && f.Mobile == false {
		//	continue
		//}

		if f.Type == GRID {
			for _, ff := range f.GetSubFields() {
				fields = append(fields, ff)
			}
		} else {
			fields = append(fields, f)
		}
	}

	if flag {
		newFileds := make([]*Field, 0)
		for _, f := range fields {
			if f.Mobile == false {
				continue
			}

			newFileds = append(newFileds, f)
		}

		return newFileds
	}

	return fields
}

func (this *Layout) AllFieldsToMap(mobile ...bool) map[string]*Field {
	fields := make(map[string]*Field)
	list := this.AllFields(mobile...)
	for _, v := range list {
		fields[v.Name] = v
	}
	return fields
}
