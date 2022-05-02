package wordlist

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

var words []string

func Initialize() error {
	data, err := ioutil.ReadFile("wordlist.txt")
	if err != nil {
		return err
	}
	words = strings.Fields(string(data))
	return nil
}

func Next() string {
	i := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(words))
	return words[i]
}
