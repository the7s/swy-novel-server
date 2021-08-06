package model

type Category struct {
	Tag  string `json:"tag"`  // 分类标签
	Name string `json:"name"` // 分类名称
	Url  string `json:"url"`  //分类url
}
