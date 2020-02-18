package math

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/jurun/workflow/utils"
)

type Float struct {
	values map[string]interface{}
}

func (this *Float) GetName() string {
	return "FLOAT"
}

func (*Float) GetCategory() string {
	return "数学函数"
}

func (*Float) GetDescription() string {
	return `FLOAT(number)
将数据转换为浮点型。`
}

func (this *Float) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Float) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return nil, errors.New("FLOAT: 参数不足")
		}

		val := utils.NewConvert(arguments[0])
		return val.Float64(), nil
	}
}
