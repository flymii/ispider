package models

import(
	"net/url"
	
	"github.com/Chain-Zhang/igo/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init(){
	dbhost := conf.AppConfig.GetString("db.host")
	dbport := conf.AppConfig.GetString("db.port")
	dbuser := conf.AppConfig.GetString("db.user")
	dbpwd := conf.AppConfig.GetString("db.password")
	dbname := conf.AppConfig.GetString("db.name")
	timezone := conf.AppConfig.GetString("db.timezone")
	if dbport == ""{
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpwd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default","mysql",dsn)
	orm.RegisterModel(new(Book), new(Chapter), new(Url), new(User), new(App))
}

func TableName(name string) string {
	return conf.AppConfig.GetString("db.prefix") + name
}