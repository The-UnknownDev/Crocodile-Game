package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/gotgbot/keyboard"

	"bot/db"
	"bot/session"
	"bot/utils"
	"bot/wordlist"
)

var commandStartHandler = handlers.NewCommand("start", commandStart)

func commandStart(b *gotgbot.Bot, ctx *ext.Context) error {
	key := fmt.Sprintf("game_%d", ctx.EffectiveChat.Id)
	if err := session.Get(key, &struct{}{}); err == nil {
		_, err = ctx.EffectiveMessage.Reply(b, "A game is already in progress.", nil)
		return err
	}
	if err := session.Set(key, &session.Game{Host: ctx.EffectiveUser.Id, Word: wordlist.Rand()}, session.GameDuration); err != nil {
		return err
	}
	if err := db.ChatsUpdate(ctx.EffectiveChat); err != nil {
		return err
	}
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("%s talks about a word.", utils.Mention(ctx.EffectiveUser.Id, ctx.EffectiveUser.FirstName)), &gotgbot.SendMessageOpts{ParseMode: "HTML", ReplyMarkup: new(keyboard.InlineKeyboard).Text("View word", "view").Row().Text("Previous word", "prev").Text("Next word", "next").Build()})
	return err
}
