package admin

import(
	"ispider/controllers"

	"strings"
)

// json 返回错误码
const (
	MSG_OK  = 0   // 成功
	MSG_ERR = -1  // 失败
)

type AdminController struct{
	controllers.BaseController
	controllerName string
	actionName string
	pageSize int
}

func (self *AdminController) Prepare(){
	self.pageSize = 10
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["cur_controller"] = self.controllerName
	self.Data["cur_action"] = self.actionName
}

func (self *AdminController) auth(){

}

func (self *AdminController) display(tpl ...string){
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/layout.html"
	self.TplName = tplname
}