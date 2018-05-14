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
type AccessLogController struct {
	beego.Controller
}

func (al *AccessLogController) Get() {

	access_log := new(models.AccessLog)

	out, _ := exec.Command("uuidgen").Output()
	id := fmt.Sprintf("%s", out[:int(len(out))-2])
	access_log.Id = id

	app := al.GetString("app")
	if app != "" {
		access_log.App = app
	}

	path := al.GetString("path")
	if path != "" {
		access_log.Path = path
	} else {
		access_log.Path = "/"
	}

	user_id := al.GetString("user_id")
	if user_id != "" {
		access_log.UserId = user_id
	}

	ip := al.GetString("ip")
	if ip != "" {
		access_log.Ip = ip
	} else {
		access_log.Ip = strings.Split(al.Ctx.Request.RemoteAddr, ":")[0]
	}

	key := al.GetString("key")
	if key != "" {
		access_log.Key = key
	}

	value := al.GetString("value")
	if value != "" {
		access_log.Value = value
	}

	meta := al.GetString("meta")
	if value != "" {
		access_log.Meta = meta
	}

	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	o.Insert(access_log)
	al.Data["json"] = access_log
	al.ServeJSON()
}
