package main

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"

	"bot/config"
	"bot/db"
	"bot/handlers"
	"bot/session"
	"bot/wordlist"
)

func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	if err := wordlist.Initialize(); err != nil {
		panic(err)
	}
	if err := session.Initialize(); err != nil {
		panic(err)
	}
	if err := db.Initialize(); err != nil {
		panic(err)
	}
	bot, err := gotgbot.NewBot(config.C.Telegram.BotToken, nil)
	if err != nil {
		panic(err)
	}
	session.Poll(bot)
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
