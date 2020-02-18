package layout

import (
	"fmt"
	"github.com/jurun/workflow/utils"
	"github.com/sirupsen/logrus"
)

type FieldType string

const (
	TEXT_FIELD             FieldType = "text"
	TEXTAREA_FIELD         FieldType = "textarea"
	NUMBER_FIELD           FieldType = "number"
	DATETIME_FIELD         FieldType = "datetime"
	DATETIME_RANGE_FIELD   FieldType = "datetime_range"
	RADIO_FIELD            FieldType = "radio"
	CHECKBOX_FIELD         FieldType = "checkbox"
	SELECT_FIELD           FieldType = "select"
	MULTI_SELECT_FIELD     FieldType = "multi_select"
	ADDRESS_FIELD          FieldType = "address"
	IDCARD_OCR_FIELD       FieldType = "idcard_ocr"
	IMAGE_FIELD            FieldType = "image"
	UPLOAD_FIELD           FieldType = "upload"
	SUBFORM_FIELD          FieldType = "subform"
	SIGNATURE_FIELD        FieldType = "signature"
	GRID                   FieldType = "grid"
	EDITOR_FIELD           FieldType = "editor"
	USER_FIELD             FieldType = "user"
	USER_MULTI_FIELD       FieldType = "multi_user"
	DEPARTMENT_FIELD       FieldType = "dept"
	DEPARTMENT_MULTI_FIELD FieldType = "multi_dept"
	SEPARATOR              FieldType = "separator"
)

func (static FieldType) String() string {
	return string(static)
}

func (static FieldType) Label() string {
	m := map[FieldType]string{
		TEXT_FIELD:             "单行文本",
		TEXTAREA_FIELD:         "多行文本",
		NUMBER_FIELD:           "数字",
		DATETIME_FIELD:         "日期时间",
		DATETIME_RANGE_FIELD:   "日期范围",
		RADIO_FIELD:            "单选框",
		CHECKBOX_FIELD:         "复选框",
		SELECT_FIELD:           "下拉框",
		MULTI_SELECT_FIELD:     "复选下拉框",
		ADDRESS_FIELD:          "地址",
		IDCARD_OCR_FIELD:       "身份证识别",
		IMAGE_FIELD:            "图片",
		UPLOAD_FIELD:           "附件",
		SUBFORM_FIELD:          "子表单",
		SIGNATURE_FIELD:        "电子签名",
		GRID:                   "栅格",
		EDITOR_FIELD:           "富文本",
		USER_FIELD:             "成员单选",
		USER_MULTI_FIELD:       "成员多选",
		DEPARTMENT_FIELD:       "部门单选",
		DEPARTMENT_MULTI_FIELD: "部门多选",
	}

	s, found := m[static]
	if found {
		return s
	}
	return ""
}

type FieldAuth struct {
	View     bool                 `json:"view"`
	Edit     bool                 `json:"edit"`
	Required bool                 `json:"required"`
	Children map[string]FieldAuth `json:"children"`
}

type Field struct {
	Name        string                 `json:"name"`
	Label       string                 `json:"label"`
	Description string                 `json:"description"`
	Type        FieldType              `json:"type"`
	View        bool                   `json:"view"`
	Edit        bool                   `json:"edit"`
	Required    bool                   `json:"required"`
	Private     bool                   `json:"privately"`
	Mobile      bool                   `json:"mobile"`
	Privacy     bool                   `json:"privacy"`
	Tab         string                 `json:"tab"`
	LineWidth   float64                `json:"line_width"`
	Setting     map[string]interface{} `json:"widget"`
}

func (this *Field) GetName() string {
	return this.Name
}

func (this *Field) GetLabel() string {
	return this.Label
}

func (this *Field) GetDescription() string {
	return this.Description
}

func (this *Field) GetType() FieldType {
	return this.Type
}

func (this *Field) GetView() bool {
	return this.View
}

func (this *Field) GetEdit() bool {
	return this.Edit
}

func (this *Field) GetRequired() bool {
	return this.Required
}

func (this *Field) GetPrivate() bool {
	return this.Private
}

func (this *Field) GetTab() string {
	return this.Label
}

func (this *Field) GetLineWidth() float64 {
	return this.LineWidth
}

func (this *Field) GetWidget() (Widget, error) {
	switch this.Type {
	case TEXT_FIELD:
		var w TextWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case TEXTAREA_FIELD:
		var w TextAreaWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case NUMBER_FIELD:
		var w NumberWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case DATETIME_FIELD:
		var w DateWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case DATETIME_RANGE_FIELD:
		var w DateRangeWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case RADIO_FIELD:
		var w RadioWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case CHECKBOX_FIELD:
		var w CheckBoxWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case SELECT_FIELD:
		var w SelectWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case MULTI_SELECT_FIELD:
		var w MultiSelectWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case ADDRESS_FIELD:
		var w AddressWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	case IDCARD_OCR_FIELD, IMAGE_FIELD, UPLOAD_FIELD:
		return nil, fmt.Errorf("不支持的组件类型: %s", this.Type.Label())
	case SUBFORM_FIELD:
		var w SubFormWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
		//case form_common.GRID:
		//	var w Grid
		//	err := utils.NewConvert(this.Setting).Bind(&w)
		//	if err != nil {
		//		return nil, err
		//	}
		//	w.Field = this
		//	return &w, nil
	case SIGNATURE_FIELD:

	case USER_FIELD:
		var w UserWidget
		err := utils.NewConvert(this.Setting).Bind(&w)
		if err != nil {
			return nil, err
		}
		w.Field = this
		return &w, nil
	}
	return nil, nil
}

func (this *Field) GetSubFields() (fields []*Field) {
	fields = make([]*Field, 0)
	if this.Type != GRID && this.Type != SUBFORM_FIELD {
		return
	}

	data, found := this.Setting["items"]
	if !found {
		logrus.Warningln("GetSubFields(), Err: Not found the items from setting")
		return
	}

	_fields := make([]*Field, 0)
	err := utils.NewConvert(data).Bind(&_fields)
	if err != nil {
		logrus.Errorln("GetSubFields(), Err: ", err.Error())
		return
	}

	this.Setting["items"] = _fields

	for _, f := range _fields {
		if f.Type == SEPARATOR {
			continue
		}
		if f.Type == GRID {
			fields = append(fields, f.GetSubFields()...)
		} else {
			fields = append(fields, f)
		}
	}

	return
}

func (this *Field) SetAuth(auth FieldAuth) {
	this.View = auth.View
	if this.View == false {
		this.Edit = false
		this.Required = false
		return
	}

	this.Edit = auth.Edit
	if this.Edit == false {
		this.Required = false
		return
	}

	this.Required = auth.Required

	// @todo 子表单 children
}
