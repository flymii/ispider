package controllers

import(
	"github.com/astaxie/beego"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct{
	beego.Controller
}

func (self *BaseController) toJson (msgno int, msg string, data interface{}){
	out := make(map[string]interface{})
	out["status"] = msgno
	out["msg"] = msg
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
}