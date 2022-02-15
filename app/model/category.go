package model

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"` // 分类名称
	Tag  string `json:"-"`    // 分类标签
	Url  string `json:"-"`    //分类url
}
