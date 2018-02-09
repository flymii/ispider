package routers

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/Chain-Zhang/igo/ilog"
	"ispider/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSCond(func(ctx *context.Context) bool {
			ilog.AppLog.Info("domain: ", ctx.Input.Domain())
			ilog.AppLog.Info("url: ",ctx.Input.URL())
			ilog.AppLog.Info("uri: ",ctx.Input.URI())
			return true
		}),
		beego.NSNamespace("/book",
			beego.NSRouter("/", &controllers.BookController{}, "get:GetAll"),
			beego.NSRouter("/getchapters", &controllers.BookController{}, "get:GetChapters"),
			beego.NSRouter("/chapter", &controllers.BookController{}, "get:GetChapter"),
		),
	)
	beego.AddNamespace(ns)
}