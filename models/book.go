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
	Status int
	From string
	Url string
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

func GetBookList(filters ...interface{})([]*Book, int64){
	books := make([]*Book, 0)
	query := orm.NewOrm().QueryTable("book")
	if len(filters) > 0{
		l := len(filters)
        for i := 0; i < l; i += 2{
			query = query.Filter(filters[i].(string), filters[i + 1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").All(&books)
	return books, total
}