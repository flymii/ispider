package routers

import(
	"github.com/astaxie/beego"
	"ispider/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/book",
			// /v1/book  获取小说列表
			beego.NSRouter("/", &controllers.BookController{}, "get:GetAll"),
			// /v1/book/getchapters?bookid=xxx&page=1  获取小说章节列表
			beego.NSRouter("/getchapters", &controllers.BookController{}, "get:GetChapters"),
			// /v1/book/chapter?id=xxx  获取章节内容
			beego.NSRouter("/chapter", &controllers.BookController{}, "get:GetChapter"),
		),
	)
	beego.AddNamespace(ns)
}