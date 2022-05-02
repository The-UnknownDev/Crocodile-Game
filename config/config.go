package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var C = struct {
	Telegram struct {
		BotToken string
	}
	Mongo struct {
		Uri string
	}
	Redis struct {
		Address  string
		Username string
		Password string
		Database int
	}
}{}

func Load() error {
	filename := os.Getenv("CONFIG")
	if filename == "" {
		filename = "config.json"
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &C)
}
