package rss

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mmcdole/gofeed"
)

type RSS struct {
	Parser *gofeed.Parser
}

func NewRSS() *RSS {
	return &RSS{
		Parser: gofeed.NewParser(),
	}
}

func (r *RSS) ParseURL(url string) (*gofeed.Feed, error) {
	return r.Parser.ParseURL(url)
}

func (r *RSS) ProcessItem(item *gofeed.Item) {
	fmt.Println("______________________________________")
	fmt.Println("Заголовок:", item.Title)
	fmt.Println("Описание:", item.Description)

	date := item.Published
	date = strings.Replace(date, "+0000", "", 1)
	fmt.Println("Дата:", date)

	content := strings.ReplaceAll(item.Content, "<p>", "\n\n")
	// Удаляем HTML-теги
	content = strings.ReplaceAll(content, "</p>", "")
	content = strings.ReplaceAll(content, "<em>", "")
	content = strings.ReplaceAll(content, "</em>", "")
	content = strings.ReplaceAll(content, "<a>", "")
	content = strings.ReplaceAll(content, "</a>", "")
	content = strings.ReplaceAll(content, "&#8220", "\"")
	content = strings.ReplaceAll(content, "&#8221", "\"")
	content = strings.ReplaceAll(content, "&#8217", "'")
	content = strings.ReplaceAll(content, "&#8216", "'")
	content = strings.ReplaceAll(content, "&amp;", "&")
	content = strings.ReplaceAll(content, "~", " ")
	content = strings.ReplaceAll(content, "…", "...")
	content = strings.ReplaceAll(content, "pic.twitter.com/", "")

	re := regexp.MustCompile(`<.*?>`)
	content = re.ReplaceAllString(content, "")
	// Удалить пустые строки
	lines := strings.Split(content, "\n")
	var result []string
	for _, line := range lines {
		if line != "" {
			result = append(result, line)
		}
	}
	content = strings.Join(result, "\n\n")
	fmt.Println("Контент:", content)

	fmt.Println("Ссылка:", item.Link)
	fmt.Println("______________________________________")
}
