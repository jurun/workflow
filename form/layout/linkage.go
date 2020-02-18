package layout

type Linkage struct {
	Condition LinkageCondition `json:"condition"`
	Field     string           `json:"field"`
	Data      LinkageData      `json:"data"`
}

type LinkageCondition struct {
	Field  string `json:"field"`
	FormId int    `json:"form_id"`
	AppId  int    `json:"app_id"`
}

type LinkageData struct {
	Field  string `json:"field"`
	FormId int    `json:"form_id"`
	AppId  int    `json:"app_id"`
}
