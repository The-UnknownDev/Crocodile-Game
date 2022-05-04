package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var C = struct {
	Mongo struct {
		Uri      string
		Database string
	}
	Redis struct {
		Address  string
		Username string
		Password string
		Database int
	}
	Telegram struct {
		BotToken string
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
