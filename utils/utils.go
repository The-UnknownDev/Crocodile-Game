package utils

import "fmt"

func Mention(id int64, text string) string {
	return fmt.Sprintf("<a href=\"tg://user?id=%d\">%s</a>", id, text)
}
