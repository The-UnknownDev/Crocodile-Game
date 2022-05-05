package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"bot/db"
	"bot/session"
	"bot/utils"
)

var commandTopPlayersHandler = handlers.NewCommand("top_players", commandTopPlayers)

func commandTopPlayers(b *gotgbot.Bot, ctx *ext.Context) error {
	topPlayers := []db.TopPlayer{}
	key := fmt.Sprintf("topplayers_%d", ctx.EffectiveChat.Id)
	err := session.Get(key, &topPlayers)
	if err != nil {
		topPlayers, err = db.TopPlayersInChat(ctx.EffectiveChat.Id)
		if err != nil {
			return err
		}
		if err = session.Set(key, topPlayers, 10*time.Minute); err != nil {
			return err
		}
	}
	text := "<b>The Top 5 Players in This Chat</b>\n\n"
	for i, p := range topPlayers {
		n := ""
		if p.Username != "" {
			n = fmt.Sprintf("@%s", p.Username)
		} else if p.FirstName != "" {
			n = p.FirstName
		} else {
			n = strconv.FormatInt(p.Id, 36)
		}
		n = utils.Mention(p.Id, n)
		s := "s"
		if p.Scores == 1 {
			s = ""
		}
		text += fmt.Sprintf("%d. %s - %d score%s\n", i+1, n, p.Scores, s)
	}
	text += "\n<i>Updates every 10 minutes.</i>"
	_, err = ctx.EffectiveMessage.Reply(b, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	return err
}
