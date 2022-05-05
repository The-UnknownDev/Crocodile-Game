package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func Load(dp *ext.Dispatcher) {
	dp.AddHandler(
		handlers.NewMessage(
			nil,
			func(b *gotgbot.Bot, ctx *ext.Context) error {
				if ctx.EffectiveChat.Type != "private" {
					return nil
				}
				return ext.ContinueGroups
			},
		),
	)
	dp.AddHandler(privateCommandTopPlayersHandler)
	dp.AddHandlerToGroup(
		handlers.NewMessage(
			nil,
			func(b *gotgbot.Bot, ctx *ext.Context) error {
				if ctx.EffectiveChat.Type != "supergroup" {
					return nil
				}
				return ext.ContinueGroups
			},
		),
		1,
	)
	dp.AddHandlerToGroup(callbackPrevHandler, 1)
	dp.AddHandlerToGroup(callbackNextHandler, 1)
	dp.AddHandlerToGroup(callbackViewHandler, 1)
	dp.AddHandlerToGroup(commandStartHandler, 1)
	dp.AddHandlerToGroup(commandTopPlayersHandler, 1)
	dp.AddHandlerToGroup(messageHandler, 1)
}
