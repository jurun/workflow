package layout

type CustomItem struct {
	Value      string   `json:"value"`
	Label      string   `json:"label"`
	ShowFields []string `json:"show_fields"`
	IsOther    bool     `json:"is_other"`
}
