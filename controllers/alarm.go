package controllers

import (
	//"crypto/tls"
	"encoding/json"
	//"github.com/astaxie/beego/httplib"
	//"strings"
	//"strconv"
	"wxapiserver/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about Alarm
type AlarmController struct {
	beego.Controller
}

type GrafanaReq struct {
	RuleId   int    `json:"ruleId"`
	RuleName string `json:"ruleName"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	State    string `json:"state"`
}

// @Title CreateAlarm_post
// @Description create a request to send alarm via POST
// @Param	content		formData  string	true		"alarm contents"
// @Success 200 send alarm success
// @Failure 403 no content
// @router /wxmsg [post]
func (u *AlarmController) Post() {
	content := u.GetString("content")
	if !models.CheckTokenActive() {
		models.GetToken()
	}
	resp, err := models.SendWxmsg(content)
	if err != nil {
		u.Ctx.WriteString(err.Error())
	} else {
		logs.Informational(resp)
		u.Ctx.WriteString("ok")
	}
}

// @Title CreateAlarmAdmin_post
// @Description create a request to send Admin alarm via POST
// @Param	content		formData  string	true		"alarm contents"
// @Success 200 send alarm success
// @Failure 403 no content
// @router /admin [post]
func (u *AlarmController) Postadmin() {
	content := u.GetString("content")
	logs.Informational("get:" + content)
	if !models.CheckTokenActive() {
		models.GetToken()
	}
	resp, err := models.SendWxmsgAdmin(content)
	if err != nil {
		u.Ctx.WriteString(err.Error())
	} else {
		logs.Informational(resp)
		u.Ctx.WriteString("ok")
	}
}

// @Title CreateAlarm_get
// @Description create a request to send alarm via GET
// @Param	content		path 	string	true		"alarm contents"
// @Success 200 send alarm success
// @Failure 403 no content
// @router /wxmsg/:content [get]
func (u *AlarmController) Get() {
	content := u.GetString(":content")
	if !models.CheckTokenActive() {
		models.GetToken()
	}
	resp, err := models.SendWxmsg(content)
	if err != nil {
		u.Ctx.WriteString(err.Error())
	} else {
		logs.Informational(resp)
		u.Ctx.WriteString("ok")
	}
}

// @Title CreateAlarm_put
// @Description alarm for grafana msg specified
// @Param	reqbody		body 	controllers.GrafanaReq	true		"body for user content.please change key to lower case"
// @Success 200 accept  and send alarm success
// @Failure 403 error occurs
// @router /grafana [put]
func (u *AlarmController) Put() {
	var rawdata string = string(u.Ctx.Input.RequestBody)
	//logs.SetLogger("console")
	logs.Informational(rawdata)
	var GrafanaReq map[string]interface{}
	json.Unmarshal([]byte(rawdata), &GrafanaReq)
	var content string = GrafanaReq["title"].(string)
	if !models.CheckTokenActive() {
		models.GetToken()
	}
	resp, err := models.SendWxmsg(content)
	if err != nil {
		u.Ctx.WriteString(err.Error())
	} else {
		logs.Informational(resp)
		u.Ctx.WriteString("ok")
	}
}

// @Title CreateAlarmCDN_put
// @Description alarm for CDN msg specified
// @Param	reqbody		body 	controllers.GrafanaReq	true		"body for user content.please change key to lower case"
// @Success 200 accept  and send alarm success
// @Failure 403 error occurs
// @router /cdn [put]
func (u *AlarmController) Putcdn() {
	var rawdata string = string(u.Ctx.Input.RequestBody)
	//logs.SetLogger("console")
	logs.Informational(rawdata)
	var GrafanaReq map[string]interface{}
	json.Unmarshal([]byte(rawdata), &GrafanaReq)
	var content string = GrafanaReq["title"].(string)
	if !models.CheckTokenActive() {
		models.GetToken()
	}
	resp, err := models.SendWxmsgCDN(content)
	if err != nil {
		u.Ctx.WriteString(err.Error())
	} else {
		logs.Informational(resp)
		u.Ctx.WriteString("ok")
	}
}

// @Title GetAllEvents
// @Description display all event
// @Success 200 get all event success
// @router /event [get]
func (u *AlarmController) GetAll() {
	events_all := models.GetAllEvents()
	var events map[string]models.Event
	events = make(map[string]models.Event)
	event := models.Event{"", "", 0, 0, ""}
	for k, v := range events_all {
		event.EventId = v.EventId
		event.EventInfo = v.EventInfo
		event.Createtime = v.Createtime
		event.Confirmtime = v.Confirmtime
		//events[k].State = v.State
		if v.State == "0" {
			event.State = "unconfirmed"
		} else if v.State == "1" {
			event.State = "manually confirmed"
		} else if v.State == "2" {
			event.State = "auto confirmed"
		}
		events[k] = event
	}
	u.Data["json"] = events
	u.ServeJSON()
}

// @Title ConfirmSpecifiedEvent
// @Description confirm event
// @Param	eventid		path 	string	true		"eventid to confirm"
// @Success 200 confirm event success
// @router /event/:eventid [get]
func (u *AlarmController) GetConfirm() {
	EventId := u.GetString(":eventid")
	resp := models.ConfirmEvent(EventId, "1")
	u.Ctx.WriteString(resp)
}

// @Title DeleteSpecifiedEvent
// @Description confirm event
// @Param	eventid		path 	string	true		"eventid to confirm"
// @Success 200 confirm event success
// @router /event/:eventid [delete]
func (u *AlarmController) DeleteEvent() {
	EventId := u.GetString(":eventid")
	resp := models.ConfirmEvent(EventId, "3")
	u.Ctx.WriteString(resp)
}
