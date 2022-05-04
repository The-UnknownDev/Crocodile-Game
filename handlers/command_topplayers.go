package handlers

import (
	"fmt"
	"html"
	"strconv"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"

	"bot/db"
	"bot/session"
)

var commandTopPlayersHandler = handlers.NewCommand("top_players", commandTopPlayers)

func commandTopPlayers(b *gotgbot.Bot, ctx *ext.Context) error {
	topPlayers, err := getTopPlayers(ctx.EffectiveChat.Id)
	if err != nil {
		return err
	}
	text := "<b>The Top 5 Players in This Chat</b>\n\n"
	for i, p := range topPlayers {
		n := ""
		if p.Username != "" {
			n = fmt.Sprintf("@%s", p.Username)
		} else if p.FirstName != "" {
			n = p.FirstName
		} else {
			n = strconv.FormatInt(p.ID, 16)
		}
		n = fmt.Sprintf("<a href=\"tg://user?id=%d\">%s</a>", p.ID, html.EscapeString(n))
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

func getTopPlayers(id int64) ([]session.TopPlayer, error) {
	topPlayers, err := session.TopPlayersGet(id)
	if err == nil {
		return topPlayers, nil
	}
	scores, err := db.ScoresTop(id)
	if err != nil {
		return nil, err
	}
	topPlayers = []session.TopPlayer{}
	for _, score := range scores {
		user, _ := db.UsersFind(score.UserID)
		topPlayers = append(topPlayers, session.TopPlayer{ID: score.UserID, Scores: score.Count, FirstName: user.FirstName, Username: user.Username})
	}
	session.TopPlayersSet(id, topPlayers)
	return topPlayers, nil
}
