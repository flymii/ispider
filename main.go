package main

import (
	"github.com/Chain-Zhang/igo/ilog"
	"github.com/Chain-Zhang/igo/conf"
	"github.com/robfig/cron"

	"ispider/spider"
	"ispider/models"
)
var i = 0

func main() {
	ilog.AppLog.Info("service start")
    c := cron.New()
	spec := conf.AppConfig.GetString("task.spec")
	ilog.AppLog.Info("spec: ",spec)
    c.AddFunc(spec,GetBook)
	c.Start()
    select{}
}

func GetBook(){
	ilog.AppLog.Info("spider start")
	books, _ := models.GetBookList("status", 1)
	for _, book := range books{
		go func(book *models.Book){
			s, err := spider.NewSpider(book.From)
			if err != nil{
				ilog.AppLog.Error("new Spider error: ", err.Error())
				return
			}
			err = s.SpiderUrl(book.Url)
			if err != nil{
				ilog.AppLog.Error("new Document error: ", err.Error())
			}
			ilog.AppLog.Info(book.Name, "已爬取完毕")
		}(book)
	}
}

