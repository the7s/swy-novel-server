package routers

type RouterGroup struct {
	UserRouter
	BookRouter
}

var RouterGroupApp = new(RouterGroup)
