package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {
	beego.GlobalControllerRouter["sheepdog/controllers:AccessLogController"] = append(beego.GlobalControllerRouter["sheepdog/controllers:AccessLogController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})
	beego.GlobalControllerRouter["sheepdog/controllers:FeedbackController"] = append(beego.GlobalControllerRouter["sheepdog/controllers:FeedbackController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})
}
