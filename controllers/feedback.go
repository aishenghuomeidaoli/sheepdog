package controllers

import (
	"os/exec"
	"fmt"
	"strings"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"sheepdog/models"
)

// Operations about Users
type FeedbackController struct {
	beego.Controller
}

func (al *FeedbackController) Get() {

	feedback := new(models.Feedback)

	out, _ := exec.Command("uuidgen").Output()
	id := fmt.Sprintf("%s", out[:int(len(out))-2])
	feedback.Id = id

	name := al.GetString("name")
	if name != "" {
		feedback.Name = name
	}

	ip := al.GetString("ip")
	if ip != "" {
		feedback.Ip = ip
	} else {
		feedback.Ip = strings.Split(al.Ctx.Request.RemoteAddr, ":")[0]
	}

	content := al.GetString("content")
	if content != "" {
		feedback.Content = content
	}

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	o.Insert(feedback)
	al.Data["json"] = feedback
	al.ServeJSON()
}
