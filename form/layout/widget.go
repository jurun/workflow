package layout

type DiffItem struct {
	Name      string      `json:"name"`
	FieldName string      `json:"field_name"`
	Type      string      `json:"type"`
	Original  interface{} `json:"original"`
	Last      interface{} `json:"last"`
}

type Diff []DiffItem

type Widget interface {
	GetValue() interface{}
	SetValue(val interface{}) error
	Diff(widget Widget) (Diff, bool)
	String() string
}

type BaseWidget struct {
	Field *Field `json:"-"`
}
