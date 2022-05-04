package handlers

import (
	"fmt"

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
	text := "The Top 5 Players in This Chat\n\n"
	for i, player := range topPlayers {
		name := player.FirstName
		if player.Username != "" {
			name = fmt.Sprintf("@%s", player.Username)
		}
		text += fmt.Sprintf("%d. %s - %d scores\n", i+1, name, player.Scores)
	}
	text += "\nThis list updates every 10 minutes."
	_, err = ctx.EffectiveMessage.Reply(b, text, nil)
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
		user, _ := db.UsersFind(score.ChatID)
		topPlayers = append(topPlayers, session.TopPlayer{ID: score.UserID, Scores: score.Count, FirstName: user.FirstName, Username: user.Username})
	}
	return topPlayers, nil
}
