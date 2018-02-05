package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Chapter struct{
	Id int
	BookId int
	Title string
	Content string
	Sort int
	Pre int
	Next int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ChapterAdd(chapter *Chapter)(int64, error){
	return orm.NewOrm().Insert(chapter)
}