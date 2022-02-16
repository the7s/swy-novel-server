package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strings"
)

// GetHtmlDoc 获取远程html内容
func GetHtmlDoc(url string) *goquery.Document {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("err : ", err)
	}

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("err : ", err)
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("err : ", err)
	}
	return doc
}

// PostHtmlDoc 获取远程html内容
func PostHtmlDoc(URL string, form url.Values) *goquery.Document {

	request, err := http.NewRequest("POST", URL, strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Println("err : ", err)
	}

	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml,application/x-www-form-urlencoded;q=0.9,image/webp,image/apng,*/*;q=0.8")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Connection", "keep-alive")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("err : ", err)
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("err : ", err)
	}
	return doc
}
