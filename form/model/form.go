package model

import (
	"fmt"
	"github.com/jurun/workflow/form/layout"
	"github.com/jurun/workflow/formula"
	"github.com/sirupsen/logrus"
	"strings"
	"time"

	"github.com/json-iterator/go"
	"github.com/jurun/workflow/utils"
)

type Form struct {
	Id           int            `xorm:"not null pk autoincr INT(11)" json:"id"`
	AppId        int            `xorm:"not null comment('应用ID') INT(11)" json:"app_id"`
	Name         string         `xorm:"not null default '' comment('表单名称') VARCHAR(200)" json:"name"`
	Description  string         `xorm:"VARCHAR(255)" json:"description"`
	IsNormal     bool           `xorm:"not null TINYINT(1)" json:"is_normal"`
	Layout       *layout.Layout `xorm:"comment('表单布局JSON（已发布并应用中的）') LONGTEXT json" json:"-"`
	LayoutChange *layout.Layout `xorm:"comment('表单布局JSON（编辑中未发布的）') LONGTEXT json" json:"-"`
	Creator      int            `xorm:"not null comment('创建人ID') INT(11)" json:"creator"`
	Created      time.Time      `xorm:"not null comment('创建日期') INT(10) CREATED" json:"-"`
	Updated      time.Time      `xorm:"not null comment('更新日期') INT(10) UPDATED" json:"-"`
	Deleted      time.Time      `xorm:"not null comment('更新日期') INT(10) DELETED" json:"-"`
	CreatedDate  int64          `xorm:"-" json:"created"`
	UpdatedDate  int64          `xorm:"-" json:"updated"`
	DeletedDate  int64          `xorm:"-" json:"deleted"`
	//LinkCode     string         `xorm:"not null comment('外链访问code') VARCHAR(200)" json:"link_code"`
	//IsPublic     bool           `xorm:"TINYINT(1) default 0" json:"is_public"`
	//LinkChannel  bool           `xorm:"TINYINT(1) default 0" json:"channel"`
	Password  string `xorm:"CHAR(32)" json:"password"`
	HasPswd   bool   `xorm:"TINYINT(1) default 0" json:"has_pswd"`
	Published bool   `xorm:"TINYINT(1) DEFAULT 0" json:"published"`
	//FormWriteAuth *Organize      `xorm:"TEXT json" json:"form_write_auth"`
	//DataReadAuth  *DataAuths     `xorm:"TEXT json" json:"data_read_auth"`
	TitleTemplate TitleTemplate `xorm:"TEXT" json:"title_template"`
	SnTemplate    SnRules       `xorm:"TEXT JSON" json:"sn_template"`
	HasChanged    bool          `xorm:"TINYINT(1) default 1" json:"has_changed"`
	Comment       bool          `xorm:"TINYINT(1) default 1" json:"comment"`
	FollowUp      bool          `xorm:"TINYINT(1) default 1" json:"follow_up"`
	//Channels      []string      `xorm:"-" json:"channels"`
	//IndexSetting  []string      `xorm:"TEXT" json:"index_setting"`
}

// 标题模板设置
type TitleTemplate string

func (static TitleTemplate) Generate(form *Form, value map[string]interface{}) (title string) {
	if string(static) == "" {
		// @todo 没取到
		fields := form.Layout.AllFields()
		if len(fields) > 0 {
			val, flag := value[fields[0].GetName()]
			if flag {
				title, _ = val.(string)
			}
		}
		return
	}

	val := make(map[string]interface{})
	fields := form.Layout.AllFieldsToMap()
	for k, v := range value {
		field, found := fields[k]
		if !found {
			val[k] = ""
			continue
		}

		wgt, err := field.GetWidget()
		if err != nil {
			val[k] = ""
			logrus.Errorln(err)
			continue
		}

		if err = wgt.SetValue(v); err != nil {
			val[k] = ""
			continue
		}

		if str := wgt.String(); str == "" {
			val[k] = ""
			continue
		}

		val[k] = wgt.String()
	}

	calculator := formula.NewCalculator(val)

	chars := strings.Split(string(static), "")
	params := make([]string, 0)

	varibleBucket := make([]string, 0)
	charBucket := make([]string, 0)
	isVariable := false
	for i := 0; i < len(chars); i++ {
		if chars[i] == "{" {
			isVariable = true
			if len(charBucket) > 0 {
				params = append(params, fmt.Sprintf("'%s'", strings.Join(charBucket, "")))
				charBucket = []string{}
			}
			continue
		}

		if chars[i] == "}" {
			isVariable = false
			if len(varibleBucket) > 0 {
				params = append(params, strings.Join(varibleBucket, ""))
				varibleBucket = []string{}
			}
			continue
		}

		if isVariable {
			varibleBucket = append(varibleBucket, chars[i])
		} else {
			charBucket = append(charBucket, chars[i])
		}

		if i == len(chars)-1 {
			if len(charBucket) > 0 {
				params = append(params, fmt.Sprintf("'%s'", strings.Join(charBucket, "")))
				charBucket = []string{}
			}
		}
	}

	args := make([]string, 0)
	for _, p := range params {
		if p == "" {
			continue
		}
		args = append(args, p)
	}
	exp := fmt.Sprintf("CONCAT(%s)", strings.Join(args, ","))

	data, err := calculator.Eval(exp)
	if err != nil {
		logrus.Errorln("Generate title failed, err:", err.Error())
		return
	}

	str, flag := data.(string)
	if !flag {
		logrus.Errorln("Generate title failed, the title is not string")
		return
	}

	title = str
	return
}

// 流水号模板设置
type SnRuleType string

const (
	INCNUMBER SnRuleType = "incNumber"
	CREATED   SnRuleType = "created"
	WIDGET    SnRuleType = "widget"
	CHARS     SnRuleType = "chars"
)

type IncNumberRule struct {
	DigitsNumber int `json:"digits_number"` // 位数，超过这个位数就从0开始
}

func (static IncNumberRule) generate(seqNum uint64) string {
	return fmt.Sprintf("%0*d", static.DigitsNumber, seqNum)
}

type CreatedRule struct {
	Format string `json:"format"`
}

func (static CreatedRule) generate(t time.Time) string {
	str := t.Format("2006-01-02-15-04-05")
	str = strings.Replace(str, "-", "", -1)
	fmt.Println("CREATED", t.String(), str)
	return str
}

type WidgetRule struct {
	Key string `json:"key"`
}

func (static WidgetRule) generate(val map[string]interface{}) string {
	v, flag := val[static.Key]
	if !flag {
		return ""
	}

	return utils.NewConvert(v).String("")
}

type CharsRule struct {
	Chars string `json:"chars"`
}

func (static CharsRule) generate() string {
	return static.Chars
}

type SnRules []struct {
	Type SnRuleType  `json:"type"`
	Rule interface{} `json:"rule"`
}

func (SnRules) Init() SnRules {
	var rules = make([]struct {
		Type SnRuleType  `json:"type"`
		Rule interface{} `json:"rule"`
	}, 0)
	return rules
}

func (static SnRules) Generate(seqNumber uint64, created time.Time, val map[string]interface{}) string {
	str := ""
	for _, rule := range static {
		b, _ := jsoniter.Marshal(rule.Rule)
		switch rule.Type {
		case INCNUMBER:
			var r IncNumberRule
			jsoniter.Unmarshal(b, &r)
			str += r.generate(seqNumber)
		case CREATED:
			var r CreatedRule
			jsoniter.Unmarshal(b, &r)
			str += r.generate(created)
		case WIDGET:
			var r WidgetRule
			jsoniter.Unmarshal(b, &r)
			str += r.generate(val)
		case CHARS:
			var r CharsRule
			jsoniter.Unmarshal(b, &r)
			str += r.generate()
		}
	}

	if str == "" {
		// @todo ...
		str = created.Local().Format("2006-01-02-15-04-05")
		str = strings.Replace(str, "-", "", -1)
	}

	return str
}
