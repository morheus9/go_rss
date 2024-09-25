package rss

import (
	"html"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
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

// Удаляет HTML-теги и заменяет специальные символы
func cleanContent(input string) string {
	// Удаляем теги <strong> и их содержимое
	re := regexp.MustCompile(`<strong>.*?</strong>`)
	input = re.ReplaceAllString(input, "")

	// Удаляем все блоки, связанные с Twitter
	reTwitterBlock := regexp.MustCompile(`<figure.*?>.*?</figure>`)
	input = reTwitterBlock.ReplaceAllString(input, "")

	// Удаляем все содержимое внутри тегов <a href="...">...</a>
	reLinks := regexp.MustCompile(`<a[^>]*>.*?</a>`)
	input = reLinks.ReplaceAllString(input, "")

	// Удаляем все блоки <figure> и их содержимое
	reFigureBlock := regexp.MustCompile(`(?s)<figure.*?>.*?</figure>`)
	input = reFigureBlock.ReplaceAllString(input, "")

	// Используем goquery для парсинга HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(input))
	if err != nil {
		return input // Возвращаем исходный текст в случае ошибки
	}

	// Получаем текст без HTML-тегов
	cleaned := doc.Find("body").Text()

	// Декодируем HTML-сущности
	cleaned = html.UnescapeString(cleaned)

	// Удалить пустые строки
	lines := strings.Split(cleaned, "\n")
	var resultContent []string
	for _, line := range lines {
		if line != "" {
			resultContent = append(resultContent, line)
		}
	}
	cleaned = strings.Join(resultContent, "\n\n")

	return cleaned
}

// Обрабатывает элемент RSS
func (r *RSS) ProcessItem(item *gofeed.Item) ItemData {
	return ItemData{
		Title:       item.Title,
		Description: item.Description,
		Date:        strings.TrimSuffix(item.Published, "+0000"),
		Content:     cleanContent(item.Content),
		Link:        item.Link,
	}
}
