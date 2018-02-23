package api


import(
	"ispider/models"
)

type BookController struct{
    ApiController
}

//获取小说列表
func (self *BookController) GetAll(){
	books, _ := models.GetBookList()
	self.ToJson(MSG_OK, "成功", books)
}

// 分页获取指定小说的章节列表, 每页10条
// url参数：bookid => 小说id; page => 页码
func (self *BookController) GetChapters(){
	bookid, err := self.GetInt("bookid")
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	page, err := self.GetInt("page")
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	chapters, _ := models.GetChapterPage(page, 10, "book_id",bookid)
	self.ToJson(MSG_OK, "success", chapters)
}

// 获取指定章节详细信息
// url参数: id => 章节id
func (self *BookController) GetChapter(){
	id, err := self.GetInt("id")
    if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	chapter, err := models.GetChapterById(id)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "success", chapter)
}