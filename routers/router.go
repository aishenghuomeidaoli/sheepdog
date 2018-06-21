package routers

import (
	"sheepdog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/access_log",
			beego.NSInclude(
				&controllers.AccessLogController{},
			),
		),
		beego.NSNamespace("/feedback",
			beego.NSInclude(
				&controllers.FeedbackController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
