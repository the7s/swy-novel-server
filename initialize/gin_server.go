package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mattn/go-colorable"
	"github.com/the7s/swy-novel-server/global"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {

	address := fmt.Sprintf(":%d", global.SWY_CONFIG.System.Addr)
	fmt.Printf(`
	欢迎使用 github.com/the7s/swy-novel-server
	当前版本:V2.0.0
	默认文档地址:http://127.0.0.1%s`, address)

	router := initGinServer()

	initRouters(router)

	s := initServer(address, router)

	s.ListenAndServe().Error()

}

func initGinServer() *gin.Engine {

	// 兼容win平台 console颜色代码块乱码的情况
	gin.ForceConsoleColor()
	gin.DefaultWriter = colorable.NewColorableStdout()

	Router := gin.Default()
	return Router
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           "127.0.0.1:8888",
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
