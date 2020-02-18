package form

import (
	"errors"
	"github.com/jurun/workflow/form/layout"
	"github.com/jurun/workflow/form/model"
)

type Form struct {
	opts Options
}

func NewForm(opt ...Option) *Form {
	opts := newOptions(opt...)
	form := &Form{
		opts: opts,
	}
	return form
}

func (this *Form) Create(appId, creator int, name string, isNormal bool, version ...int) (*model.Form, error) {
	m := model.Form{
		AppId:        appId,
		Name:         name,
		Creator:      creator,
		IsNormal:     isNormal,
		Layout:       layout.Layout{}.Init(),
		LayoutChange: layout.Layout{}.Init(),
		SnTemplate:   model.SnRules{}.Init(),
	}

	_, err := this.opts.storage.Insert(&m)
	if err != nil {
		return &model.Form{}, err
	}

	Form{}.afterSelect(&m)
	return &m, nil
}

func (this *Form) Get(id int, unscoped ...bool) (*model.Form, bool, error) {
	_unscoped := false
	if len(unscoped) > 0 {
		_unscoped = unscoped[0]
	}

	var (
		m     model.Form
		found bool
		err   error
	)

	if _unscoped {
		found, err = this.opts.storage.ID(id).Unscoped().After(Form{}.afterSelect).Get(&m)
	} else {
		found, err = this.opts.storage.ID(id).After(Form{}.afterSelect).Get(&m)
	}

	if err != nil || !found {
		return &model.Form{}, found, err
	}

	return &m, true, nil
}

func (this *Form) Update(form *model.Form) error {
	var err error
	form.HasChanged = true

	_, err = this.opts.storage.
		MustCols("is_normal", "is_public",
			"has_changed", "published",
			"comment", "follow_up",
		).
		ID(form.Id).Update(form)

	if err != nil {
		return err
	}

	Form{}.afterSelect(form)
	return nil
}

func (this *Form) Delete(id int) error {

	fm, found, err := this.Get(id)
	if err != nil {
		return err
	}

	if !found {
		return nil
	}

	if fm.Published == true {
		return errors.New("当前表单启用中，请禁用后再删除")
	}

	session := this.opts.storage.NewSession()
	session.Begin()

	defer session.Close()
	_, err = session.ID(id).Delete(new(model.Form))
	if err != nil {
		session.Rollback()
		return err
	}

	if err = this.opts.process.DeleteAll(id, session); err != nil {
		session.Rollback()
		return err
	}

	session.Commit()
	return nil
}

func (this *Form) List(ids []int, unscope ...bool) ([]*model.Form, map[int]*model.Form, error) {
	var c bool
	if len(unscope) > 0 {
		c = unscope[0]
	}

	var err error
	forms := make([]*model.Form, 0)
	mp := make(map[int]*model.Form)

	if c == false {
		err = this.opts.storage.In("id", ids).After(Form{}.afterSelect).Find(&forms)
	} else {
		err = this.opts.storage.In("id", ids).Unscoped().After(Form{}.afterSelect).Find(&forms)
	}
	if err != nil {
		return nil, nil, err
	}

	for _, f := range forms {
		mp[f.Id] = f
	}

	return forms, mp, err
}

func (static Form) afterSelect(bean interface{}) {
	fm := bean.(*model.Form)
	fm.CreatedDate = fm.Created.Unix()
	fm.UpdatedDate = fm.Updated.Unix()
	fm.DeletedDate = fm.Deleted.Unix()
	if fm.DeletedDate < 0 {
		fm.DeletedDate = 0
	}

	//if fm.FormWriteAuth == nil {
	//    fm.FormWriteAuth = Organize{}.Init()
	//}
	//if fm.DataReadAuth == nil {
	//    fm.DataReadAuth = DataAuths{}.Init()
	//}

	//if fm.Channels == nil {
	//    fm.Channels = make([]string, 0)
	//}

	//if fm.IndexSetting == nil {
	//    fm.IndexSetting = make([]string, 0)
	//}
}
