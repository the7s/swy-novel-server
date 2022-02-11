package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/the7s/swy-novel-server/app/service"
	"github.com/the7s/swy-novel-server/library/utils"
	"net/http"
)

var Category = new(CategoryApi)

type CategoryApi struct{}

func (category CategoryApi) GetList(c *gin.Context) {

	data := service.Category.GetCategories("")
	c.JSON(http.StatusOK, data)
}

func (category CategoryApi) GetBookList(c *gin.Context) {
	categoryTag := c.Param("categoryTag")

	categoryUrl := "" + utils.SwyDecodeUrl(categoryTag)

	data := service.Book.GetBooks(categoryUrl)
	c.JSON(http.StatusOK, data)

}
