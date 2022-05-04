package session

import (
	"bot/db"
	"encoding/json"
	"fmt"
	"time"
)

const topPlayersDuration = 10 * time.Minute

func topPlayersKey(id int64) string {
	return fmt.Sprintf("topplayers_%d", id)
}

func TopPlayersSet(id int64, topPlayers []db.TopPlayer) error {
	key := topPlayersKey(id)
	value, err := json.Marshal(topPlayers)
	if err != nil {
		return err
	}
	_, err = client.Set(client.Context(), key, value, topPlayersDuration).Result()
	return err
}

func TopPlayersGet(id int64) ([]db.TopPlayer, error) {
	key := topPlayersKey(id)
	value, err := client.Get(client.Context(), key).Result()
	if err != nil {
		return nil, err
	}
	topPlayers := []db.TopPlayer{}
	return topPlayers, json.Unmarshal([]byte(value), &topPlayers)
}
