package admin

import(
	"ispider/models"

	"github.com/Chain-Zhang/igo/util"
	"github.com/astaxie/beego/utils/pagination"

	"time"
)

type UserController struct{
	AdminController
}

func (self *UserController)Users(){
	page, err := self.GetInt("p")
	if err != nil{
		page = 1
	}
	users, total := models.GetUserList(page, self.pageSize, "status__gte", 0)
	self.Data["users"] = users
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

func (self *UserController)AjaxAdd(){
	username := self.GetString("username")
	if len(username) < 1{
		self.ToJson(MSG_ERR, "The username can not be empty", nil)
	}
	user := new(models.User)
	user.Username = username
	user.Password = util.Md5("123456", false)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Level = 99
	user.Status = 0
	id, err := models.UserAdd(user)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	if id < 1{
		self.ToJson(MSG_ERR, "新增用户失败，换个用户名再试", nil)
	}
	self.ToJson(MSG_OK, "新增用户成功", nil)
}

func (self *UserController)AjaxDelete(){
	id, err := self.GetInt("id")
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user, err := models.GetUserById(id)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user.Status = -1
	user.UpdatedAt = time.Now()
	err = user.Update()
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "用户【"+user.Username+"】已删除", nil)
}