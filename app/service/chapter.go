package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
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

		var chapter = model.Chapter{
			Name:    name,
			Tag:     utils.SwyEncodeUrl(tag),
			Content: "",
			HasNext: true,
		}
		chapterList = append(chapterList, chapter)
	})
	return chapterList
}

func (cs ChapterService) GetDetail(webUrl string) model.ChapterDetail {
	fmt.Println(webUrl)

	var doc = utils.GetHtmlDoc(webUrl)

	a := doc.Find("#wrapper .content_read .box_con")

	content, _ := a.Find("#content").Html()
	var chapter = model.Chapter{
		Name:    a.Find(".bookname h1").Text(),
		Tag:     "",
		Content: content,
		HasNext: true,
	}

	var chapterDetail = model.ChapterDetail{
		CurrentChapter: chapter,
		PreChapterUrl:  "",
		NextChapterUrl: "",
	}

	return chapterDetail
}
