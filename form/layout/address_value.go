package layout

import "fmt"

type AddressValue struct {
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	Detail   string `json:"detail"`
}

func (AddressValue) Init() *AddressValue {
	return &AddressValue{}
}

func (static AddressValue) String() string {
	return fmt.Sprintf("%s%s%s%s", static.Province, static.City, static.District, static.Detail)
}
