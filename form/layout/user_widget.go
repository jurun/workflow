package layout

import (
	"github.com/jurun/workflow/utils"
)

type userWidget_userList struct {
	Deps        []int `json:"deps"`
	Roles       []int `json:"roles"`
	Users       []int `json:"users"`
	CurrentUser bool  `json:"current_user"`
}

type UserWidget struct {
	BaseWidget
	Limit   *userWidget_userList `json:"limit"`
	Value   int                  `json:"value"`
	Linkage Linkage              `json:"linkage"`
}

func (this *UserWidget) GetValue() interface{} {
	return &utils.Convert{}
}

func (this *UserWidget) SetValue(val interface{}) error {
	return nil
}

func (this *UserWidget) Diff(widget Widget) (Diff, bool) {
	return Diff{}, false
}

func (this *UserWidget) String() string {
	return ""
}
