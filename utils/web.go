package utils

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"strings"
)

func GetWebText(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP请求失败:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应错误:", err)
		return "", err
	}
	text := ""
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	doc.Find("*").Each(func(i int, s *goquery.Selection) {
		node := s.Nodes[0].Data
		currentText := ""
		if node == "p" {
			currentText = s.Text() + "\n"
		} else if node == "pre" {
			currentText = s.Text() + "\n"
		} else if node == "li" {
			currentText = s.Text() + "\n"
		}
		if len(currentText) > 15 {
			text += currentText
		}
	})
	return text, nil
}
