package handlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"bot/game"
	"bot/wordlist"
)

var callbackNextHandler = handlers.NewCallback(
	func(cq *gotgbot.CallbackQuery) bool {
		return cq.Data == "next"
	},
	callbackNext,
)

func callbackNext(b *gotgbot.Bot, ctx *ext.Context) error {
	game, err := game.Get(ctx.EffectiveChat.Id)
	if err != nil {
		return err
	}
	if game.Host != ctx.EffectiveUser.Id {
		_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: "This is not for you."})
	} else {
		game.Word = wordlist.Next(game.Word)
		if err = game.Set(); err != nil {
			return err
		}
		_, err = ctx.CallbackQuery.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: wordlist.Get(game.Word)})
	}
	return err
}
