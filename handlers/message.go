package handlers

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"bot/db"
	"bot/session"
	"bot/wordlist"
)

var messageHandler = handlers.NewMessage(
	func(msg *gotgbot.Message) bool {
		return msg.Text != ""
	},
	message,
)

func message(b *gotgbot.Bot, ctx *ext.Context) error {
	game := session.Game{}
	key := fmt.Sprintf("game_%d", ctx.EffectiveChat.Id)
	err := session.Get(key, &game)
	if err != nil {
		return err
	}
	if ctx.EffectiveUser.Id == game.Host {
		return err
	}
	if strings.Contains(strings.ToLower(ctx.EffectiveMessage.Text), wordlist.Get(game.Word)) {
		if err = session.Del(key); err != nil {
			return err
		}
		if err = db.UsersUpdate(ctx.EffectiveUser); err != nil {
			return err
		}
		if err = db.ScoresUpdate(ctx.EffectiveChat.Id, ctx.EffectiveUser.Id); err != nil {
			return err
		}
		_, err = ctx.EffectiveMessage.Reply(b, fmt.Sprintf("%s guessed the word correctly!", ctx.EffectiveUser.FirstName), nil)
	}
	return err
}
