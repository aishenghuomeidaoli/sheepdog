package models

import (
	"github.com/astaxie/beego/orm"
)

type Feedback struct {
	Id      string `orm:"pk"`
	Ip      string `orm:"null"`
	Name    string `orm:"null"`
	Content string `orm:"null"`
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Feedback))
}
