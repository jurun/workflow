package form

import (
	"github.com/go-xorm/xorm"
	"github.com/jurun/workflow/process"
)

type Option func(options *Options)

type Options struct {
	storage *xorm.Engine
	process *process.Process
	debug   bool
}

func Storage(engine *xorm.Engine) Option {
	return func(options *Options) {
		options.storage = engine
	}
}

func Process(process *process.Process) Option {
	return func(options *Options) {
		options.process = process
	}
}

func Debug(flag bool) Option {
	return func(options *Options) {
		options.debug = flag
	}
}

func newOptions(opts ...Option) Options {
	opt := Options{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}
