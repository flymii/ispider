package controllers

import(
	"ispider/models"
)

type BookController struct{
    BaseController
}

func (self *BookController) GetAll(){
	books, _ := models.GetBookList()
	self.toJson(MSG_OK, "成功", books)
}

func (self *BookController) GetChapters(){
	bookid, err := self.GetInt("bookid")
	if err != nil{
		self.toJson(MSG_ERR, err.Error(), nil)
	}
	page, err := self.GetInt("page")
	if err != nil{
		self.toJson(MSG_ERR, err.Error(), nil)
	}
	chapters, _ := models.GetChapterPage(page, 10, "book_id",bookid)
	self.toJson(MSG_OK, "success", chapters)
}

func (self *BookController) GetChapter(){
	id, err := self.GetInt("id")
    if err != nil{
		self.toJson(MSG_ERR, err.Error(), nil)
	}
	chapter, err := models.GetChapterById(id)
	if err != nil{
		self.toJson(MSG_ERR, err.Error(), nil)
	}
	self.toJson(MSG_OK, "success", chapter)
}