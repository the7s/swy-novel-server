package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/the7s/swy-novel-server/app/service"
	"github.com/the7s/swy-novel-server/library/utils"
	"net/http"
)

type bookApi struct{}

var Book = new(bookApi)

func (b bookApi) GetIndex(c *gin.Context) {
	webUrl := ""
	bl := service.Book.GetBooks(webUrl)
	c.JSON(http.StatusOK, bl)
}

func (b bookApi) GetDetail(c *gin.Context) {

	webUrl := "" + utils.SwyDecodeUrl(c.Param("bookTag"))
	data := service.Book.GetBookDetail(webUrl)
	c.JSON(http.StatusOK, data)
}

func (b bookApi) GetChapter(c *gin.Context) {

	webUrl := "" + utils.SwyDecodeUrl(c.Param("bookTag"))
	data := service.Chapter.GetChapterList(webUrl)
	c.JSON(http.StatusOK, data)
}

func (b bookApi) GetChapterDetail(c *gin.Context) {

	webUrl := "" + utils.SwyDecodeUrl(c.Param("chapterTag"))
	data := service.Chapter.GetDetail(webUrl, c.Param("bookTag"))
	c.JSON(http.StatusOK, data)
}
