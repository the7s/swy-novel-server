package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/the7s/swy-novel-server/routers"
)

// 初始化总路由

func initRouters(r *gin.Engine) {

	// 健康监测
	PublicGroup := r.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}

	routerGroupApp := routers.RouterGroupApp

	routerGroupApp.UserRouter.InitUserRouter(PublicGroup)
	routerGroupApp.BookRouter.InitBookRouter(PublicGroup)

}
