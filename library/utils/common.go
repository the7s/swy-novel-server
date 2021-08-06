package utils

import "strings"

func SwyDecodeUrl(url string) string {

	return strings.ReplaceAll(url, "^", "/")
}

func SwyEncodeUrl(url string) string {

	return strings.ReplaceAll(url, "/", "^")
}

func SwyParseColon(str string) string {
	return str[(strings.Index(str, "：") + 3):]
}
