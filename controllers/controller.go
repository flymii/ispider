package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct {
	beego.Controller
}

// 固定返回的json数据格式
// msgno: 错误码
// msg: 错误信息
// data: 返回数据
func (self *BaseController) ToJson(msgno int, msg string, data interface{}) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["msg"] = msg
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
}

//获取用户IP地址
func (self *BaseController) GetClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
