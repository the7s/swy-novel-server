package model

type Chapter struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Content string `json:"content"`
	HasNext bool   `json:"hasNext"`
}

type ChapterDetail struct {
	CurrentChapter Chapter
	PreChapterUrl  string `json:"preChapterUrl"`
	NextChapterUrl string `json:"nextChapterUrl"`
}
