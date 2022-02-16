package service

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/the7s/swy-novel-server/app/model"
	"github.com/the7s/swy-novel-server/global"
	"github.com/the7s/swy-novel-server/library/utils"
	"net/url"
	"strings"
	"sync"
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

func (bs bookService) SearchBookDetail(sBookName string, sAuthor string) model.BookDetail {
	searchPageUrl := global.SWY_CONFIG.Website.BQGUrl + "modules/article/waps.php"

	form := url.Values{}
	form.Add("searchkey", sBookName)

	var doc = utils.PostHtmlDoc(searchPageUrl, form)
	var bookDetail = model.BookDetail{}
	var isBreak bool
	doc.Find("#checkform tbody tr").Each(func(i int, s *goquery.Selection) {
		// 跳过首行
		if i > 0 && !isBreak {
			name := s.Find("td").Eq(0).Find("a").Text()
			novelUrl, _ := s.Find("td").Eq(0).Find("a").Attr("href")
			author := s.Find("td").Eq(2).Text()
			newChapter := s.Find("td").Eq(1).Find("a").Text()
			newChapterUrl, _ := s.Find("td").Eq(1).Find("a").Attr("href")

			//if name == sBookName && author == sAuthor {
			if name == sBookName {
				isBreak = true
				bookDetail.Url = novelUrl
				bookDetail.Name = name
				bookDetail.Author = author
				bookDetail.NewChapter = newChapter
				bookDetail.NewChapterUrl = global.SWY_CONFIG.Website.BQGUrl + newChapterUrl[1:]
				bs.completeBookDetail(&bookDetail)
			}
		}
	})
	return bookDetail
}

func (bs bookService) SearchBook(sBookName string) []*model.Book {
	var wg sync.WaitGroup
	searchPageUrl := global.SWY_CONFIG.Website.BQGUrl + "modules/article/waps.php"
	form := url.Values{}
	form.Add("searchkey", sBookName)
	var doc = utils.PostHtmlDoc(searchPageUrl, form)
	var bl []*model.Book

	doc.Find("#checkform tbody tr").Each(func(i int, s *goquery.Selection) {
		// 跳过首行
		if i > 0 {
			name := s.Find("td").Eq(0).Find("a").Text()
			novelUrl, _ := s.Find("td").Eq(0).Find("a").Attr("href")
			author := s.Find("td").Eq(2).Text()
			var book = &model.Book{
				Name:     name,
				Url:      novelUrl,
				CoverUrl: "",
				Author:   author,
				Desc:     "",
				Status:   "",
			}
			bl = append(bl, book)
		}
	})
	wg.Add(len(bl))
	for _, b := range bl {

		func(book *model.Book) {
			defer wg.Done()
			var doc = utils.GetHtmlDoc(book.Url)
			var pEle = doc.Find("#wrapper .box_con")
			desc := pEle.Find("#maininfo #intro p").Eq(1).Text()
			coverUrl, _ := pEle.Find("#sidebar #fmimg img").Attr("src")
			book.Desc = desc
			book.CoverUrl = coverUrl

		}(b)
	}
	wg.Wait()
	return bl
}

func (bs bookService) completeBookDetail(bookDetail *model.BookDetail) {

	var doc = utils.GetHtmlDoc(bookDetail.Url)
	var pEle = doc.Find("#wrapper .box_con")

	lastUpdateAt := pEle.Find("#maininfo #info p").Eq(2).Text()
	desc := pEle.Find("#maininfo #intro p").Eq(1).Text()
	coverUrl, _ := pEle.Find("#sidebar #fmimg img").Attr("src")

	bookDetail.Desc = desc
	bookDetail.LastUpdateAt = strings.Replace(lastUpdateAt, "最后更新：", "", 1)
	bookDetail.CoverUrl = coverUrl

	chapters := make([]model.Chapter, doc.Find(".box_con #list dl dd").Size())

	doc.Find(".box_con #list dl dd").Each(func(i int, s *goquery.Selection) {
		chapterName := s.Find("a").Text()
		chapterUrl, _ := s.Find("a").Attr("href")
		chapters[i] = model.Chapter{
			Name:    chapterName,
			Url:     global.SWY_CONFIG.Website.BQGUrl + chapterUrl[1:],
			HasNext: true,
		}
	})
	chapters[len(chapters)-1].HasNext = false
	bookDetail.Chapters = chapters
}

func (bs bookService) GetChapterDetail(chapterUrl string) model.ChapterDetail {
	var doc = utils.GetHtmlDoc(chapterUrl)
	doc.Find(".content_read #content p").Remove()
	content := doc.Find(".content_read #content").Text()
	content = strings.ReplaceAll(content, "    ", "<p>")
	content = strings.ReplaceAll(content, "\n\n", "</p>")
	content += "</p>"

	preChapterUrl, _ := doc.Find(".bottem2 a").Eq(1).Attr("href")
	nextChapterUrl, _ := doc.Find(".bottem2 a").Eq(3).Attr("href")

	preChapterUrl = global.SWY_CONFIG.Website.BQGUrl + preChapterUrl[1:]
	nextChapterUrl = global.SWY_CONFIG.Website.BQGUrl + nextChapterUrl[1:]

	var haxNext = true
	if !strings.Contains(preChapterUrl, ".html") {
		preChapterUrl = ""
	}

	if !strings.Contains(nextChapterUrl, ".html") {
		haxNext = false
		nextChapterUrl = ""
	}

	return model.ChapterDetail{
		CurrentChapter: model.Chapter{
			Name:    doc.Find(".bookname h1").Text(),
			Url:     chapterUrl,
			Content: content,
			HasNext: haxNext,
		},
		PreChapterUrl:  preChapterUrl,
		NextChapterUrl: nextChapterUrl,
	}
}
