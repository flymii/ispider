package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct{
	Id int
	Username string
	Password string
	Level int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (self *User)TableName()string{
	return "user"
}

func UserAdd(user *User)(int64, error){
	return orm.NewOrm().Insert(user)
}
