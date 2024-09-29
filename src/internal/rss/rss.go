package rss

import (
	"html"
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
	Title          string
	Description    string
	Date           string
	Content        string
	FirstParagraph string
	Link           string
}

// Удаляет HTML-теги и заменяет специальные символы, возвращает только первый непустой параграф
func cleanContent(input string) string {
	// Используем goquery для парсинга HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(input))
	if err != nil {
		return "" // Возвращаем пустую строку в случае ошибки
	}

	// Получаем текст без HTML-тегов
	cleaned := doc.Find("body").Text()
	cleaned = html.UnescapeString(cleaned)

	// Разбиваем текст на параграфы и удаляем пустые строки
	paragraphs := strings.Split(cleaned, "\n")
	for _, paragraph := range paragraphs {
		trimmed := strings.TrimSpace(paragraph)
		if trimmed != "" {
			return trimmed // Возвращаем первый непустой параграф
		}
	}

	return "" // Если нет непустых параграфов, возвращаем пустую строку
}

// Удаляет пустые строки из текста
func cleanDescription(input string) string {
	// Разбиваем текст на строки и удаляем пустые строки
	lines := strings.Split(input, "\n")
	var cleanedLines []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			cleanedLines = append(cleanedLines, trimmed) // Добавляем непустую строку
		}
	}
	return strings.Join(cleanedLines, "\n") // Объединяем непустые строки обратно в текст
}

// Обрабатывает элемент RSS
func (r *RSS) ProcessItem(item *gofeed.Item) ItemData {
	cleanedContent := cleanContent(item.Content) // Очищаем контент и получаем первый параграф
	cleanedDescription := cleanDescription(item.Description)
	return ItemData{
		Date:        strings.TrimSuffix(item.Published, "+0000"),
		Title:       item.Title,
		Description: cleanedDescription,
		Content:     cleanedContent,
		Link:        item.Link,
	}
}
