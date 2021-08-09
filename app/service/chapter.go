package service

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
	"swy-novel-server/app/model"
	"swy-novel-server/library/utils"
)

type ChapterService struct{}

var Chapter = ChapterService{}

func (cs ChapterService) GetChapterList(webUrl string) []model.Chapter {

	var chapterList []model.Chapter
	var doc = utils.GetHtmlDoc(webUrl)

	a := doc.Find("#wrapper .box_con #list dl").Children()

	var dtNum int

	a.Each(func(i int, s *goquery.Selection) {

		if s.Get(0).Data == "dt" {
			dtNum++
		}
		if dtNum <= 1 || s.Get(0).Data == "dt" {
			return
		}

		name := s.Find("a").Text()
		tag, _ := s.Find("a").Attr("href")

		var hasNext = true

		if i+1 == a.Length() {
			hasNext = false
		}

		var chapter = model.Chapter{
			Name:    name,
			Tag:     utils.SwyEncodeUrl(tag),
			Content: "",
			HasNext: hasNext,
		}
		chapterList = append(chapterList, chapter)
	})
	return chapterList
}

func (cs ChapterService) GetDetail(webUrl string, bookTag string) model.ChapterDetail {

	var doc = utils.GetHtmlDoc(webUrl)

	a := doc.Find("#wrapper .content_read .box_con")

	pre, _ := a.Find(".bottem2 a").Eq(0).Attr("href")
	next, _ := a.Find(".bottem2 a").Eq(2).Attr("href")

	var hasNext bool
	if strings.Contains(next, "html") {
		hasNext = true
	}

	content, _ := a.Find("#content").Html()
	var chapter = model.Chapter{
		Name:    a.Find(".bookname h1").Text(),
		Tag:     bookTag,
		Content: content,
		HasNext: hasNext,
	}

	var chapterDetail = model.ChapterDetail{
		CurrentChapter: chapter,
		PreChapterUrl:  utils.SwyEncodeUrl(pre),
		NextChapterUrl: utils.SwyEncodeUrl(next),
	}

	return chapterDetail
}
