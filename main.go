package main

import (
	"github.com/the7s/swy-novel-server/global"
	"github.com/the7s/swy-novel-server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://gpproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod dowload

func main() {
	global.SWY_VP = initialize.Viper()
	global.SWY_LOG = initialize.Zap()
	initialize.RunServer()
}
