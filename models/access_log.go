package models

import (
	"github.com/astaxie/beego/orm"
)

type AccessLog struct {
	Id     string `orm:"pk"`
	App    string `orm:"null"`
	Path   string `orm:"null"`
	UserId string `orm:"null"`
	Ip     string `orm:"null"`
	Key    string `orm:"null"`
	Value  string `orm:"null"`
	Meta   string `orm:"null"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(AccessLog))
}
