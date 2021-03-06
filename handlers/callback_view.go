package handlers

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"bot/session"
	"bot/wordlist"
)

var callbackViewHandler = handlers.NewCallback(
	func(cq *gotgbot.CallbackQuery) bool {
		return cq.Data == "view"
	},
	callbackView,
)

func callbackView(b *gotgbot.Bot, ctx *ext.Context) error {
	game := session.Game{}
	err := session.Get(fmt.Sprintf("game_%d", ctx.EffectiveChat.Id), &game)
	if err != nil {
		return err
	}
	if game.Host != ctx.EffectiveUser.Id {
		_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: "This is not for you."})
	} else {
		_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: wordlist.Get(game.Word)})
	}
	return err
}
