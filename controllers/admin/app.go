package admin

import(
	"time"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/Chain-Zhang/igo/util"
	"ispider/models"
)

type AppController struct{
	AdminController
}

func (self *AppController)Apps(){
	self.display()
}

func (self *AppController) MyApps(){
	page, err := self.GetInt("p")
	if err != nil{
		page = 1
	}
	apps, total := models.GetAppList(page, self.pageSize, "status__gte", 0, "user_id", self.login_userId)
	self.Data["apps"] = apps
	self.Data["total"] = total
	if total < 1{
		self.Data["hasdata"] = false
	}else{
		self.Data["hasdata"] = true
	}

	paginator := pagination.SetPaginator(self.Ctx, self.pageSize, total)
	self.Data["paginator"] = paginator
	self.display()
}

func (self *AppController)Add(){
	self.display()
}

func (self *AppController)AjaxAdd(){
	appname := self.GetString("appname")
	desc := self.GetString("desc")

	app := new(models.App)
	app.Appname = appname
	app.Desc = desc
	app.Count = 1000
	app.CreatedAt = time.Now()
	app.Status = 0
	app.Token = util.Md5(appname + app.CreatedAt.String(), false)
	app.UserId = self.login_userId

	_, err := models.AppAdd(app)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "新增应用成功", nil)
}