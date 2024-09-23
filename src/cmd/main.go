package main

import (
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
		rss.ProcessItem(item)
	}
}
