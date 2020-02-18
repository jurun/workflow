package model

import (
	"time"
)

type Process struct {
	Id          int       `xorm:"not null pk autoincr INT(11)" json:"id"`
	FormId      int       `xorm:"not null INT(11)" json:"form_id"`
	Description string    `xorm:"VARCHAR(255)" json:"description"`
	WithDraw    bool      `xorm:"not null TINYINT(1) default 0" json:"with_draw"`
	AutoSubmit  bool      `xorm:"TINYINT(1)" json:"auto_submit"`
	State       bool      `xorm:"not null TINYINT(1) default 0" json:"state"`
	Nodes       *Nodes    `xorm:"not null TEXT json" json:"-"`
	NodesChange *Nodes    `xorm:"not null TEXT json" json:"-"`
	Version     int       `xorm:"not null INT(11)" json:"version"`
	Created     time.Time `xorm:"not null INT(10) CREATED" json:"-"`
	CreatedAt   string    `xorm:"-" json:"created"`
	Updated     time.Time `xorm:"not null INT(10) UPDATED" json:"-"`
	UpdatedAt   string    `xorm:"-" json:"updated"`
	Deleted     time.Time `xorm:"not null INT(10) DELETED" json:"deleted"`
}
