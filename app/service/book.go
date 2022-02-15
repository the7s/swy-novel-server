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
