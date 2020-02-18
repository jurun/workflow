package process

import (
	"github.com/go-xorm/xorm"
)

type Option func(options *Options)

type Options struct {
	storage *xorm.Engine
	debug   bool
}

func Storage(engine *xorm.Engine) Option {
	return func(options *Options) {
		options.storage = engine
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
