package v1

type ApiGroup struct {
	UserApi
	BookApi
}

var ApiGroupApp = new(ApiGroup)
