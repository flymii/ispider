package main

import (
	"ispider/spider"
    "fmt"
	"github.com/Chain-Zhang/igo/ilog"
)

func main() {
	ilog.AppLog.Info("start")
	s, err := spider.NewSpider("booktxt")
	if err != nil{
		ilog.AppLog.Fatal("new Spider error: ", err.Error())
	}
	err = s.SpiderUrl("http://www.booktxt.net/2_2219/")
	if err != nil{
		ilog.AppLog.Fatal("new Document error: ", err.Error())
	}
	var str string
	fmt.Scan(&str)
}

