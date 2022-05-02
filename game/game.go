package game

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"

	"bot/config"
	"bot/redis"
)

const GameDuration = 5 * time.Minute

type Game struct {
	Chat int64
	Host int64
	Word int
}

func Poll(bot *gotgbot.Bot) {
	ch := redis.Client.Subscribe(redis.Client.Context(), fmt.Sprintf("__keyevent@%d__:expired", config.C.Redis.Database)).Channel()
	go func() {
		message := <-ch
		go func() {
			chatId, err := strconv.ParseInt(message.Payload, 10, 64)
			if err != nil {
				return
			}
			bot.SendMessage(chatId, "Time is up!", nil)
		}()
	}()
}

func (game *Game) Set() error {
	key := strconv.FormatInt(game.Chat, 10)
	value, err := json.Marshal(game)
	if err != nil {
		return err
	}
	_, err = redis.Client.Set(redis.Client.Context(), key, value, GameDuration).Result()
	return err
}

func (game *Game) Del() (bool, error) {
	key := strconv.FormatInt(game.Chat, 10)
	i, err := redis.Client.Del(redis.Client.Context(), key).Result()
	return i > 0, err
}

func Get(chat int64) (*Game, error) {
	game := &Game{Chat: chat}
	key := strconv.FormatInt(chat, 10)
	value, err := redis.Client.Get(redis.Client.Context(), key).Result()
	if err != nil {
		return game, err
	}
	err = json.Unmarshal([]byte(value), game)
	return game, err
}
