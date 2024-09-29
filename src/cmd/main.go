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
	rss := rss.NewRSS()

	// Парсим RSS
	feed, err := rss.ParseURL(cfg.FeedURL)
	if err != nil {
		log.Fatal(err)
	}

	// Обрабатываем элементы RSS
	for _, item := range feed.Items {
		data := rss.ProcessItem(item)
		fmt.Printf("DATE: %s\n", data.Date)
		fmt.Printf("TITLE: %s\n", data.Title)
		fmt.Printf("DESCRIPTION: %s\n", data.Description)
		fmt.Printf("CONTENT: %s\n", data.Content)
		fmt.Printf("LINK: %s\n", data.Link)
	}
}
