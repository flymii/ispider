package controllers

import(
	"github.com/astaxie/beego"
)

// json 返回错误码
const (
	MSG_OK  = 0   // 成功
	MSG_ERR = -1  // 失败
)

// 基类
type BaseController struct{
	beego.Controller
}

// 固定返回的json数据格式
// msgno: 错误码
// msg: 错误信息
// data: 返回数据
func (self *BaseController) toJson (msgno int, msg string, data interface{}){
	out := make(map[string]interface{})
	out["status"] = msgno
	out["msg"] = msg
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
}