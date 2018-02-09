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

func GetChapterPage(page, pageSize int, filters ...interface{})([]*Chapter, int64){
	offset := (page - 1) * pageSize
	list := make([]*Chapter, 0)
	query := orm.NewOrm().QueryTable("chapter")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("sort").Limit(pageSize, offset).All(&list)
	return list, total
}

func GetChapterById(id int)(*Chapter, error){
	chapter := new(Chapter)
	err := orm.NewOrm().QueryTable("chapter").Filter("id", id).One(chapter)
	if err != nil{
		return nil, err
	}
	return chapter, nil
}