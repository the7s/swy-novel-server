package service

import (
	"github.com/PuerkitoBio/goquery"
	"swy-novel-server/app/model"
	"swy-novel-server/library/utils"
)

type CategoryService struct{}

var Category = CategoryService{}

func (cs CategoryService) GetCategories(webUrl string) []model.Category {
	var categoryList []model.Category

	var doc = utils.GetHtmlDoc(webUrl)

	// Find the review items
	doc.Find("#wrapper .nav ul li").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		if i == 0 || i == 7 || i == 14 || i == 15 {
			return
		}
		name := s.Find("a").Text()
		tag, _ := s.Find("a").Attr("href")

		var category = model.Category{
			Tag:  utils.SwyEncodeUrl(tag),
			Name: name,
			Url:  webUrl + tag,
		}
		categoryList = append(categoryList, category)
	})
	return categoryList
}
