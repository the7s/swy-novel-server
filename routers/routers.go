package routers

import (
	"github.com/gin-gonic/gin"
	"swy-novel-server/app/api"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()

	r.GET("/", api.Index.HelloWorld)

	bookR := r.Group("/book")
	{
		bookR.GET("/", api.Book.GetIndex)
		bookR.GET("/:bookTag", api.Book.GetDetail)
		bookR.GET("/:bookTag/chapter", api.Book.GetChapter)
		bookR.GET("/:bookTag/chapter/:chapterTag", api.Book.GetChapterDetail)
	}

	categoryR := r.Group("/category")
	{
		categoryR.GET("/", api.Category.GetList)
		categoryR.GET("/:categoryTag/books", api.Category.GetBookList)
	}

	return r
}
