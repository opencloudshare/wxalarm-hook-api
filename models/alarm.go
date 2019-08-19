package models

import (
	"errors"
	//"strconv"
	"time"

	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"strings"
)

var (
	GetTokenUrl string
	Token       WxToken
	EventMap    map[string]*Event
)

func init() {
	logs.Async()
	logs.SetLogger(logs.AdapterFile, `{"filename":"/root/go/src/wxapiserver/logs/BackendService.log","level":7}`)
	GetTokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=xxx&corpsecret=xxx"
	Token := WxToken{"", 0, false}
	EventMap = make(map[string]*Event)
	e := Event{"0", "no biggie", time.Now().Unix(), 0, "0"}
	EventMap["0"] = &e
	if !CheckTokenActive() {
		GetToken()
	}
	fmt.Println(Token.Accesstoken)
}

type Event struct {
	EventId     string
	EventInfo   string
	Createtime  int64
	Confirmtime int64
	State       string
}

type WxToken struct {
	Accesstoken string
	Createtime  int64
	Active      bool
}

type SendWxData struct {
	Toparty string
	Msgtype string
	Agentid int
	Safe    int
	Text    DataText
}

type DataText struct {
	Content string
}

type GrafanaReq struct {
	RuleId   int    `json:ruleId`
	RuleName string `json:ruleName`
	Title    string `json:"title"`
	Message  string `json:"message"`
	State    string `json:"state"`
}

func GetToken() (err error) {
	req := httplib.Get(GetTokenUrl)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	str, err := req.String()
	var response map[string]interface{}
	json.Unmarshal([]byte(str), &response)
	if err != nil {
		logs.Error(err.Error())
		return
	}
	if response["access_token"] == nil {
		logs.Error("no access token got, check weixin service")
		return
	}
	Token = WxToken{response["access_token"].(string), time.Now().Unix(), true}
	logs.Informational("get access token ok")
	logs.Informational(response["access_token"].(string))
	return
}

func CheckTokenActive() (s bool) {
	if Token.Createtime+7000 > time.Now().Unix() {
		return true
	} else {
		return false
	}
}

func SendWxmsg(c string) (s string, e error) {
	var SendWxmsgUrl string = strings.Join([]string{"https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=", Token.Accesstoken}, "")
	//data := SendWxData{"2", "text", 1000002, 0, DataText{c}}
	data := SendWxData{"4", "text", 1000002, 0, DataText{c}}
	data_j, err := json.Marshal(data)
	if err != nil {
		e = errors.New("convert data to json error, check  service")
		logs.Error(err.Error())
		return "", e
	}
	req2 := httplib.Post(SendWxmsgUrl)
	req2.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req2.Body(strings.ToLower(string(data_j)))
	s, e = req2.String()
	if e != nil {
		e = errors.New("send req to wxapi error, check weixin service")
		logs.Error(err.Error())
		return "", e
	}
	return s, nil
}

func SendWxmsgCDN(c string) (s string, e error) {
	var SendWxmsgUrl string = strings.Join([]string{"https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=", Token.Accesstoken}, "")
	//data := SendWxData{"2", "text", 1000002, 0, DataText{c}}
	data := SendWxData{"3", "text", 1000002, 0, DataText{c}}
	data_j, err := json.Marshal(data)
	if err != nil {
		e = errors.New("convert data to json error, check  service")
		logs.Error(err.Error())
		return "", e
	}
	req2 := httplib.Post(SendWxmsgUrl)
	req2.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req2.Body(strings.ToLower(string(data_j)))
	s, e = req2.String()
	if e != nil {
		e = errors.New("send req to wxapi error, check weixin service")
		logs.Error(err.Error())
		return "", e
	}
	return s, nil
}

func SendWxmsgAdmin(c string) (s string, e error) {
	var SendWxmsgUrl string = strings.Join([]string{"https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=", Token.Accesstoken}, "")
	//data := SendWxData{"2", "text", 1000002, 0, DataText{c}}
	data := SendWxData{"2", "text", 1000002, 0, DataText{c}}
	data_j, err := json.Marshal(data)
	if err != nil {
		e = errors.New("convert data to json error, check  service")
		logs.Error(err.Error())
		return "", e
	}
	req2 := httplib.Post(SendWxmsgUrl)
	req2.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req2.Body(strings.ToLower(string(data_j)))
	s, e = req2.String()
	if e != nil {
		e = errors.New("send req to wxapi error, check weixin service")
		logs.Error(err.Error())
		return "", e
	}
	return s, nil
}

func GetAllEvents() map[string]*Event {
	return EventMap
}

func CreateEvent(k string, i string) {
	e := Event{k, i, time.Now().Unix(), 0, "0"}
	EventMap[k] = &e
	return
}

func ConfirmEvent(e string, r string) (s string) {
	if v, ok := EventMap[e]; ok {
		if v.State == "1" {
			s = "manually confirmed"
		} else if v.State == "2" {
			s = "seems auto confirmed earlier, ignored"
		} else if v.State == "3" {
			s = "deleted"
		} else if v.State == "0" {
			EventMap[e].State = r
			EventMap[e].Confirmtime = time.Now().Unix()
			if r == "1" || r == "2" {
				s = "confirm ok"
			} else if r == "3" {
				s = "delete ok"
			} else {
				s = "other state"
			}
		}
	} else {
		s = "no event for this id found"
	}
	return s
}
