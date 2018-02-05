package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type Book struct{
	Id int
	Name string
	Author string
	Image string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func BookAdd(book *Book)(int64, error){
	return orm.NewOrm().Insert(book)
}

func GetBookByName(name string)(*Book, error){
	book := new(Book)
	err := orm.NewOrm().QueryTable("book").Filter("name", name).One(book)
	if err != nil || book.Id < 1{
		return nil, err
	}
	return book, nil
}