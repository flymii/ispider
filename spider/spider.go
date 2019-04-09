package spider

import (
	"errors"

	"github.com/Chain-Zhang/igo/conf"
	"github.com/Chain-Zhang/igo/ilog"
	"github.com/robfig/cron"

	"ispider/models"
)

type SBook struct {
	Name     string
	Image    string
	Url      string
	Chapters []*SChapter
}

type SChapter struct {
	Title   string
	Url     string
	Order   int
	Pre     int
	Next    int
	Content string
}

type Spider interface {
	SpiderUrl(url string) error
}

func NewSpider(from string) (Spider, error) {
	switch from {
	case "booktxt":
		return new(BookTextSpider), nil
	default:
		return nil, errors.New("系统暂未处理该类型的配置文件")
	}
}

func Start() {
	ilog.Info("service start")
	c := cron.New()
	spec := conf.AppConfig.GetString("task.spec")
	ilog.Info("spec: ", spec)
	c.AddFunc(spec, getBook)
	c.Start()
	select {}
}

func getBook() {
	ilog.Info("spider start")
	books, _ := models.GetBookList("status", 1)
	for _, book := range books {
		go func(book *models.Book) {
			s, err := NewSpider(book.From)
			if err != nil {
				ilog.Error("new Spider error: ", err.Error())
				return
			}
			err = s.SpiderUrl(book.Url)
			if err != nil {
				ilog.Error("new Document error: ", err.Error())
			}
			ilog.Info(book.Name, "已爬取完毕")
		}(book)
	}
}
