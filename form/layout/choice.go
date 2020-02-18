package layout

type Choice struct {
	BaseWidget
	CustomItems   []CustomItem  `json:"custom_items"`
	RemoteItems   RemoteItems   `json:"remote_items"`
	LinkFormItems LinkFormItems `json:"link_form_items"`
}
