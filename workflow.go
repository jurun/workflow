package workflow

import (
	"github.com/jurun/workflow/form"
	"github.com/jurun/workflow/process"
)

type Workflow struct {
	Options Options
	Form    *form.Form
	Process *process.Process
	//Instance *instance.Instance
}

func NewWorkflow(option ...Option) (*Workflow, error) {
	opt, err := newOptions(option...)
	if err != nil {
		return &Workflow{}, err
	}

	opt.Form = form.NewForm(
		form.Storage(opt.Mysql),
		form.Debug(opt.config.debug),
		form.Process(opt.Process),
	)

	opt.Process = process.NewProcess(
		process.Storage(opt.Mysql),
		process.Debug(opt.config.debug),
	)

	wkf := &Workflow{
		Form:    opt.Form,
		Process: opt.Process,
		Options: opt,
	}

	return wkf, nil
}
