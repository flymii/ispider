package main

import (
	_ "ispider/routers"
	"ispider/spider"
	"github.com/astaxie/beego"
	"github.com/Chain-Zhang/igo/ilog"
)

func main() {
	ilog.Info("start")
	go spider.Start()
	beego.Run(":8089")
}


