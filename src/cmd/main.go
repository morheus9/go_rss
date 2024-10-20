package main

import (
	"fmt"
	"log"

	"github.com/morheus9/go_rss/src/internal/config"
	"github.com/morheus9/go_rss/src/internal/rss"
)

func main() {
	// Загрузка конфига
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatal(err)
	}

	// Создаем новый парсер RSS
	rssParser := rss.NewRSS()

	// Парсим RSS
	// Обрабатываем каждый URL в FeedURL
	for _, url := range cfg.FeedURL {
		feed, err := rssParser.ParseURL(url)
		if err != nil {
			log.Printf("Error parsing URL %s: %v\n", url, err)
			continue
		}

		// Обрабатываем элементы RSS
		for _, item := range feed.Items {
			data, err := rssParser.ProcessItem(item)
			if err != nil {
				log.Printf("Error processing item: %v\n", err)
				continue
			}
			fmt.Printf("DATE: %s\n", data.Date)
			fmt.Printf("TITLE: %s\n", data.Title)
			fmt.Printf("DESCRIPTION: %s\n", data.Description)
			fmt.Printf("LINK: %s\n", data.Link)
		}
	}
}
