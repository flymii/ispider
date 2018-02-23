package admin

type LoginController struct{
	AdminController
}

func (self *LoginController) Login(){
	self.TplName = "login/login.html"
}