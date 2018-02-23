package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type App struct{
	Id int
	Appname string
	Token string
	UserId int
	CreatedAt time.Time
}

func (self *App)TableName()string{
	return "app"
}

func AppAdd(app *App)(int64, error){
	return orm.NewOrm().Insert(app)
}
