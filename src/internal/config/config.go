package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	FeedURL []string `json:"FeedURL"`
}

func LoadConfig(filename string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
