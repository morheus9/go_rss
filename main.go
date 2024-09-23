package main

import (
	"log"

	"src/rss"

	"src/config"
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
