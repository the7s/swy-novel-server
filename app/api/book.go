package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swy-novel-server/app/service"
	"swy-novel-server/config"
	"swy-novel-server/library/utils"
)

type bookApi struct{}

var Book = new(bookApi)

func (b bookApi) GetIndex(c *gin.Context) {
	webUrl := config.Config.GetWebUrl()
	bl := service.Book.GetBooks(webUrl)
	c.JSON(http.StatusOK, bl)
}

func (b bookApi) GetDetail(c *gin.Context) {

	webUrl := config.Config.GetWebUrl() + utils.SwyDecodeUrl(c.Param("bookTag"))
	data := service.Book.GetBookDetail(webUrl)
	c.JSON(http.StatusOK, data)
}

func (b bookApi) GetChapter(c *gin.Context) {

	webUrl := config.Config.GetWebUrl() + utils.SwyDecodeUrl(c.Param("bookTag"))
	data := service.Chapter.GetChapterList(webUrl)
	c.JSON(http.StatusOK, data)
}

func (b bookApi) GetChapterDetail(c *gin.Context) {

	webUrl := config.Config.GetWebUrl() + utils.SwyDecodeUrl(c.Param("chapterTag"))
	data := service.Chapter.GetDetail(webUrl)
	c.JSON(http.StatusOK, data)
}
