package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type indexApi struct{}

var Index = new(indexApi)

func (index *indexApi) HelloWorld(c *gin.Context) {

	c.JSON(http.StatusOK, "Hi,欢迎使用 SWY小说 API")
}
