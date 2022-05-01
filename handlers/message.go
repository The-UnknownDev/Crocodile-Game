package handlers

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"bot/game"
)

var messageHandler = handlers.NewMessage(
	func(msg *gotgbot.Message) bool {
		return msg.Text != ""
	},
	message,
)

func message(b *gotgbot.Bot, ctx *ext.Context) error {
	game, err := game.Get(ctx.EffectiveChat.Id)
	if err != nil {
		return err
	}
	if ctx.EffectiveUser.Id == game.Host {
		return err
	}
	if strings.Contains(strings.ToLower(ctx.EffectiveMessage.Text), game.Word) {
		if _, err = game.Del(); err != nil {
			return err
		}
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("%s guessed the word correctly!", ctx.EffectiveUser.FirstName), nil)
	}
	return err
}
