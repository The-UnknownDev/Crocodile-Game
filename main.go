package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"bot/config"
	"bot/game"
	"bot/handlers"
	"bot/redis"
	"bot/wordlist"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	if err := wordlist.Initialize(); err != nil {
		panic(err)
	}
	if err := redis.Initialize(); err != nil {
		panic(err)
	}
	bot, err := gotgbot.NewBot(config.C.Telegram.BotToken, nil)
	if err != nil {
		panic(err)
	}
	game.Poll(bot)
	updater := ext.NewUpdater(nil)
	handlers.Load(updater.Dispatcher)
	if err = updater.StartPolling(
		bot, &ext.PollingOpts{
			DropPendingUpdates: true,
		},
	); err != nil {
		panic(err)
	}
	updater.Idle()
}
