package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgbot/keyboard"

	"bot/game"
	"bot/wordlist"
)

var commandStartHandler = handlers.NewCommand("start", commandStart)

func commandStart(b *gotgbot.Bot, ctx *ext.Context) error {
	if _, err := game.Get(ctx.EffectiveChat.Id); err == nil {
		_, err = ctx.EffectiveMessage.Reply(b, "A game is already in progress.", nil)
		return err
	}
	if err := (&game.Game{Chat: ctx.EffectiveChat.Id, Host: ctx.EffectiveUser.Id, Word: wordlist.Rand()}).Set(); err != nil {
		return err
	}
	_, err := ctx.EffectiveMessage.Reply(
		b,
		fmt.Sprintf("%s talks about a word", ctx.EffectiveUser.FirstName),
		&gotgbot.SendMessageOpts{
			ReplyMarkup: new(keyboard.InlineKeyboard).Text("View word", "view").Row().Text("Previous word", "prev").Text("Next word", "next").Build(),
		},
	)
	return err
}
