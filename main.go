package main

import (
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/joho/godotenv"

	"bot/game"
	"bot/handlers"
	"bot/wordlist"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	if err := wordlist.Initialize(); err != nil {
		panic(err)
	}
	if err := game.Initialize(); err != nil {
		panic(err)
	}
	token := os.Getenv("BOT_TOKEN")
	bot, err := gotgbot.NewBot(token, nil)
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
