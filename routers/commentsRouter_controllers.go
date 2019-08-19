package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"] = append(beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"],
		beego.ControllerComments{
			Method: "Postadmin",
			Router: `/admin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"] = append(beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"],
		beego.ControllerComments{
			Method: "Putcdn",
			Router: `/cdn`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"] = append(beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/event`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"] = append(beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"],
		beego.ControllerComments{
			Method: "GetConfirm",
			Router: `/event/:eventid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"] = append(beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"],
		beego.ControllerComments{
			Method: "DeleteEvent",
			Router: `/event/:eventid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"] = append(beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/grafana`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"] = append(beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/wxmsg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"] = append(beego.GlobalControllerRouter["wxapiserver/controllers:AlarmController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/wxmsg/:content`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
