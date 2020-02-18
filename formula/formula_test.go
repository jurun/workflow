package formula

import (
	"testing"
)

func init() {
}

func TestLogic(t *testing.T) {
	str := `OR(field1>90,AND(field2>100,field3=='向志'))`
	formula := NewCalculator(map[string]interface{}{
		"field1": 80,
		"field2": 801,
		"field3": "向志",
	})

	ret, err := formula.Eval(str)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", ret)
}
