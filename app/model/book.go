package model

type Book struct {
	Name     string `json:"name"`        // 名称
	CoverUrl string `json:"coverUrl"`    // 封面
	Url      string `json:"url"`         // 书籍地址
	Author   string `json:"author"`      // 作者
	Desc     string `json:"description"` // 描述
	Status   string `json:"status"`      // 小说状态
}

type BookList []Book

type BookDetail struct {
	Name          string    `json:"name"`          // 名称
	Url           string    `json:"url"`           // 书地址
	CoverUrl      string    `json:"coverUrl"`      // 封面
	Author        string    `json:"author"`        // 作者
	Desc          string    `json:"description"`   // 描述
	Status        string    `json:"status"`        //状态
	LastUpdateAt  string    `json:"lastUpdateAt"`  //最后更新时间
	NewChapter    string    `json:"newChapter"`    //最新一章、
	NewChapterUrl string    `json:"newChapterUrl"` //最新一章地址
	Chapters      []Chapter `json:"chapters"`      // 章节目录
}
