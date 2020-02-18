package layout

import (
	"github.com/jurun/workflow/utils"
)

type UserMultiWidget struct {
	BaseWidget
	Limit   *userWidget_userList `json:"limit"`
	Value   []int                `json:"value"`
	Linkage Linkage              `json:"linkage"`
}

func (this *UserMultiWidget) GetValue() interface{} {
	return &utils.Convert{}
}

func (this *UserMultiWidget) SetValue(val interface{}) error {
	return nil
}

func (this *UserMultiWidget) Diff(widget Widget) (Diff, bool) {
	return Diff{}, false
}

func (this *UserMultiWidget) String() string {
	return ""
}
