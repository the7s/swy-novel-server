package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"swy-novel-server/app/service"
	"swy-novel-server/config"
	"swy-novel-server/library/utils"
)

var Category = new(CategoryApi)

type CategoryApi struct{}

func (category CategoryApi) GetList(c *gin.Context) {

	data := service.Category.GetCategories(config.Config.GetWebUrl())
	c.JSON(http.StatusOK, data)
}

func (category CategoryApi) GetBookList(c *gin.Context) {
	categoryTag := c.Param("categoryTag")

	categoryUrl := config.Config.GetWebUrl() + utils.SwyDecodeUrl(categoryTag)

	data := service.Book.GetBooks(categoryUrl)
	c.JSON(http.StatusOK, data)

}
