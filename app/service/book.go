package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/the7s/swy-novel-server/app/model"
	"github.com/the7s/swy-novel-server/library/utils"
)

type bookService struct{}

var Book = new(bookService)

func (bs bookService) GetQDBooks(webUrl string) []model.Book {
	var doc = utils.GetHtmlDoc(webUrl)
	var bl []model.Book
	// Find the review items
	doc.Find(".main-content-wrap .all-book-list ul li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title

		url, _ := s.Find(".book-img-box a").Attr("href")
		coverUrl, _ := s.Find(".book-img-box a img").Attr("src")
		name := s.Find(".book-mid-info h2 a").Text()
		author := s.Find(".book-mid-info .author a.name").Text()
		desc := s.Find(".book-mid-info p.intro").Text()
		status := s.Find(".book-mid-info .author span").Text()

		var book = model.Book{
			Name:     name,
			CoverUrl: "https:" + coverUrl,
			Url:      "https:" + url,
			Author:   author,
			Desc:     desc,
			Status:   status,
		}
		bl = append(bl, book)
	})
	return bl
}

func (bs bookService) GetBooks(webUrl string) []model.Book {

	var doc = utils.GetHtmlDoc(webUrl)

	var bl []model.Book

	// Find the review items
	doc.Find("#main #hotcontent div .item").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		author := s.Find("dl dt span").Text()
		name := s.Find("dl dt a").Text()
		url, _ := s.Find("dl dt a").Attr("href")
		coverUrl, _ := s.Find(".image a img").Attr("src")
		desc := s.Find("dl dd").Text()

		var book = model.Book{
			Tag:      utils.SwyEncodeUrl(url),
			Name:     name,
			CoverUrl: "" + coverUrl,
			Url:      "" + url,
			Author:   author,
			Desc:     desc,
		}
		bl = append(bl, book)
	})
	return bl
}

func (bs bookService) GetBookDetail(webUrl string) model.BookDetail {

	var doc = utils.GetHtmlDoc(webUrl)

	var pEle = doc.Find("#wrapper .box_con")

	name := pEle.Find("#maininfo #info h1").Text()

	author := pEle.Find("#maininfo #info p").Eq(0).Text()

	status := pEle.Find("#maininfo #info p").Eq(1).Text()

	lastUpdateAt := pEle.Find("#maininfo #info p").Eq(2).Text()

	newChapter := pEle.Find("#maininfo #info p").Eq(3).Find("a").Text()
	newChapterUrl, _ := pEle.Find("#maininfo #info p").Eq(3).Find("a").Attr("href")

	desc := pEle.Find("#maininfo #intro p").Eq(0).Text()
	coverUrl, _ := pEle.Find("#sidebar #fmimg img").Attr("src")

	var bookDetail = model.BookDetail{
		Name:          name,
		CoverUrl:      "" + coverUrl,
		Author:        utils.SwyParseColon(author),
		Desc:          desc,
		Status:        utils.SwyParseColon(status),
		LastUpdateAt:  utils.SwyParseColon(lastUpdateAt),
		NewChapter:    newChapter,
		NewChapterUrl: utils.SwyEncodeUrl(newChapterUrl),
	}
	return bookDetail
}
