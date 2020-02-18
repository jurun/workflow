package process

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/jurun/workflow/process/model"
)

type Process struct {
	opts Options
}

func NewProcess(opt ...Option) *Process {
	opts := newOptions(opt...)
	process := &Process{
		opts: opts,
	}
	return process
}

func (this *Process) Create(formId int) (*model.Process, error) {
	version, err := this.GetMaxVersion(formId)
	if err != nil {
		return &model.Process{}, err
	}
	version++

	nodes := model.Nodes{}.Init()
	nodes.AddNode(model.START_NODE)
	p := model.Process{
		FormId:      formId,
		Nodes:       nodes,
		NodesChange: nodes,
		Version:     version,
	}

	_, err = this.opts.storage.InsertOne(&p)
	if err != nil {
		return &model.Process{}, err
	}
	return &p, nil
}

func (this *Process) Get(id int) (*model.Process, bool, error) {
	var p model.Process
	found, err := this.opts.storage.ID(id).Get(&p)
	if err != nil || !found {
		return &model.Process{}, found, err
	}

	return &p, true, nil
}

func (this *Process) List(formId int, onlyEnabled bool) ([]*model.Process, error) {
	datas := make([]*model.Process, 0)
	var err error
	if onlyEnabled {
		err = this.opts.storage.
			Where("form_id=? and state = 1", formId).
			Desc("version").
			Iterate(new(model.Process), func(idx int, bean interface{}) error {
				attr := bean.(*model.Process)
				datas = append(datas, attr)
				return nil
			})
	} else {
		err = this.opts.storage.
			Where("form_id=?", formId).
			Desc("version").
			Iterate(new(model.Process), func(idx int, bean interface{}) error {
				attr := bean.(*model.Process)
				datas = append(datas, attr)
				return nil
			})
	}

	return datas, err
}

func (this *Process) GetMaxVersion(formId int) (int, error) {
	var md model.Process
	found, err := this.opts.storage.Where("form_id=?", formId).Unscoped().Desc("version").Get(&md)
	if err != nil {
		return 0, err
	}
	if !found {
		return 0, nil
	}

	return md.Version, nil
}

func (this *Process) Update(process *model.Process) error {
	_, err := this.opts.storage.ID(process.Id).MustCols("with_draw", "auto_submit", "state").Update(process)
	return err
}

func (this *Process) Enable(id int) error {
	ps, found, err := this.Get(id)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("要启用的流程数据不存在")
	}

	session := this.opts.storage.NewSession()
	session.Begin()
	defer session.Begin()

	_, err = session.Where("form_id=?", ps.FormId).
		NotIn("id", []int{ps.Id}).
		Cols("state").
		MustCols("state").
		Unscoped().
		Update(&model.Process{
			State: false,
		})
	if err != nil {
		session.Rollback()
		return err
	}

	ps.Nodes = ps.NodesChange
	ps.State = true
	_, err = session.ID(ps.Id).Cols("state", "nodes").MustCols("state").Unscoped().Update(ps)
	if err != nil {
		session.Rollback()
		return err
	}

	session.Commit()
	return nil
}

func (this *Process) Delete(id int, session ...*xorm.Session) error {

	ps, found, err := this.Get(id)
	if err != nil {
		return err
	}
	if !found {
		return nil
	}

	if ps.State == true {
		return errors.New("当前流程启用中，请禁用后再删除")
	}

	var ss *xorm.Session
	if len(session) > 0 {
		ss = session[0]
	} else {
		ss = this.opts.storage.NewSession()
	}

	_, err = ss.ID(id).Delete(&model.Process{})
	if err != nil {
		return fmt.Errorf("删除流程失败, Err: %s", err.Error())
	}
	return nil
}

func (this *Process) DeleteAll(formId int, session *xorm.Session) error {
	ids := make([]int, 0)

	err := session.Where("form_id=?", formId).Iterate(new(model.Process), func(idx int, bean interface{}) error {
		ids = append(ids, bean.(*model.Process).Id)
		return nil
	})

	if err != nil {
		return err
	}

	if len(ids) == 0 {
		return nil
	}

	_, err = session.In("id", ids).Delete(new(model.Process))
	return err
}
