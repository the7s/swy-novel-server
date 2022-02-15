package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/the7s/swy-novel-server/app/construct"
	"github.com/the7s/swy-novel-server/app/model"
	"github.com/the7s/swy-novel-server/app/service"
	"github.com/the7s/swy-novel-server/global"
	"net/http"
	"strconv"
)

type BookApi struct{}

func (b BookApi) GetBook(c *gin.Context) {
	var CategoryConst = construct.CategoryConst
	categoryStr := c.DefaultQuery("categoryId", "0")
	pageIdStr := c.DefaultQuery("pageId", "1")
	pageId, _ := strconv.Atoi(pageIdStr)
	categoryId, _ := strconv.Atoi(categoryStr)

	category, ok := CategoryConst[categoryId]
	if !ok {
		c.JSON(http.StatusOK, "error")
	}
	var categoryTag = category.Tag

	if pageId > 1 {
		switch categoryId {
		case 0:
			categoryTag = category.Tag + "/" + "page" + pageIdStr
		default:
			categoryTag = category.Tag + "-page" + pageIdStr
		}
	}
	webUrl := global.SWY_CONFIG.Website.QiDianUrl + categoryTag
	bl := service.Book.GetQDBooks(webUrl)
	c.JSON(http.StatusOK, bl)
}

func (b BookApi) GetCategory(c *gin.Context) {
	var categories [16]model.Category
	var CategoryConst = construct.CategoryConst
	for k, v := range CategoryConst {
		categories[k] = v
	}
	c.JSON(http.StatusOK, categories)
}

func (b BookApi) Search(c *gin.Context) {
	var categories [16]model.Category
	var CategoryConst = construct.CategoryConst
	for k, v := range CategoryConst {
		categories[k] = v
	}
	c.JSON(http.StatusOK, categories)
}
