package wordlist

import (
	"io/ioutil"
	"math/rand"
	"strings"
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
	return words[rand.Intn(len(words))]
}
