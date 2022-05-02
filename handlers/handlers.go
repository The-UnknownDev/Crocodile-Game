package handlers

import (

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Load(dp *ext.Dispatcher) {
	dp.AddHandler(
		handlers.NewMessage(
			func(msg *gotgbot.Message) bool {
				return msg.Chat.Type != "supergroup" && msg.Chat.Type != "group"
			},
			func(b *gotgbot.Bot, ctx *ext.Context) error {
				return ext.EndGroups
			},
		),
	)
	dp.AddHandler(callbackNextHandler)
	dp.AddHandler(callbackViewHandler)
	dp.AddHandler(commandStartHandler)
	dp.AddHandler(messageHandler)
}