package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/the7s/swy-novel-server/app/api/v1"
)

type BookRouter struct{}

func (br BookRouter) InitBookRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	bookApi := v1.ApiGroupApp.BookApi

	bookRouter := Router.Group("v1/books")
	{
		bookRouter.GET("", bookApi.GetBook)
		bookRouter.GET("category", bookApi.GetCategory)
		bookRouter.GET("search", bookApi.Search)
		bookRouter.GET("searchAll", bookApi.SearchAll)
		bookRouter.GET("chapter", bookApi.GetChapterDetail)
	}
	return bookRouter
}
