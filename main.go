package main

import (
	"github.com/Chain-Zhang/igo/ilog"
	"github.com/astaxie/beego"
	_ "ispider/routers"
	"ispider/spider"
)

func main() {
	ilog.Info("start")
	go spider.Start()
	beego.Run(":8089")
}
