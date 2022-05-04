package session

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"

	"bot/config"
)

const gameDuration = 5 * time.Minute

type Game struct {
	Host int64
	Word int
}


func gameKey(id int64) string {
	return fmt.Sprintf("game_%d", id)
}

func gamePoll(bot *gotgbot.Bot) {
	ch := client.Subscribe(client.Context(), fmt.Sprintf("__keyevent@%d__:expired", config.C.Redis.Database)).Channel()
	go func() {
		message := <-ch
		go func() {
			slices := strings.Split(message.Payload, "_")
			if slices[0] != "game" {
				return
			}
			id, err := strconv.ParseInt(slices[1], 10, 64)
			if err != nil {
				return
			}
			bot.SendMessage(id, "Time is up!", nil)
		}()
	}()
}

func GameSet(id int64, game *Game) error {
	key := gameKey(id)
	value, err := json.Marshal(game)
	if err != nil {
		return err
	}
	_, err = client.Set(client.Context(), key, value, gameDuration).Result()
	return err
}

func GameDel(id int64) error {
	key := gameKey(id)
	_, err := client.Del(client.Context(), key).Result()
	return err
}

func GameGet(id int64) (*Game, error) {
	key := gameKey(id)
	value, err := client.Get(client.Context(), key).Result()
	if err != nil {
		return nil, err
	}
	game := &Game{}
	return game, json.Unmarshal([]byte(value), game)
}
