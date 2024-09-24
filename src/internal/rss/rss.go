package rss

import (
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

type ItemData struct {
	Title       string
	Description string
	Date        string
	Content     string
	Link        string
}

func (r *RSS) ProcessItem(item *gofeed.Item) ItemData {
	var data ItemData
	data.Title = item.Title
	data.Description = item.Description
	data.Date = strings.TrimSuffix(item.Published, "+0000")
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
	var resultContent []string
	for _, line := range lines {
		if line != "" {
			resultContent = append(resultContent, line)
		}
	}
	content = strings.Join(resultContent, "\n\n")
	data.Content = content
	data.Link = item.Link
	return data
}
