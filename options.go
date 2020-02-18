package workflow

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/jurun/workflow/form"
	"github.com/jurun/workflow/instance"
	"github.com/jurun/workflow/process"
	"gopkg.in/mgo.v2"
)

type Option func(options *Options)

type Options struct {
	Mysql *xorm.Engine
	Mongo *mgo.Database

	Form     *form.Form
	Process  *process.Process
	Instance *instance.Instance

	config struct {
		mysqlAddr, mysqlUser, mysqlPswd, mysqlDbName, mongoDns, mongoDbname string
		debug                                                               bool
	}
}

func newMysql(user, pswd, host, dbName string) (*xorm.Engine, error) {
	if user == "" {

	}
	if host == "" {

	}

	if dbName == "" {

	}

	u := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&interpolateParams=true&parseTime=true&loc=Local",
		user,
		pswd,
		host,
		dbName,
	)

	engine, err := xorm.NewEngine("mysql", u)
	return engine, err
}

func newMongo(dns, dbName string) (*mgo.Database, error) {
	session, err := mgo.Dial(dns)
	if err != nil {
		return &mgo.Database{}, err
	}

	session.SetMode(mgo.Monotonic, true)
	s := session.Copy()
	return s.DB(dbName), nil
}

func newOptions(opts ...Option) (Options, error) {
	opt := Options{}
	for _, o := range opts {
		o(&opt)
	}

	var err error
	opt.Mysql, err = newMysql(
		opt.config.mysqlUser,
		opt.config.mysqlPswd,
		opt.config.mysqlAddr,
		opt.config.mysqlDbName,
	)
	if err != nil {
		return Options{}, err
	}

	opt.Mongo, err = newMongo(opt.config.mongoDns, opt.config.mongoDbname)
	if err != nil {
		return Options{}, err
	}
	return opt, nil
}

func Debug() Option {
	return func(options *Options) {
		options.config.debug = true
	}
}

func Mysql(host, user, pswd, dbName string) Option {

	return func(options *Options) {
		options.config.mysqlAddr = host
		options.config.mysqlUser = user
		options.config.mysqlPswd = pswd
		options.config.mysqlDbName = dbName
	}
}

func MongoDB(dns, dbName string) Option {
	return func(options *Options) {
		options.config.mongoDns = dns
		options.config.mongoDbname = dbName
	}
}
