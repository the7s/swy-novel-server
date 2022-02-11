package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/the7s/swy-novel-server/app/api/v1"
)

type UserRouter struct{}

func (ur UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	userApi := v1.ApiGroupApp.UserApi
	{
		//userRouter.POST("register", userApi.Register) // 注册用户
		userRouter.POST("login", userApi.Login) // 注册用户
	}

	return userRouter
}
