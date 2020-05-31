package code

import (
	"strings"
)

func Encode(str string) string {
	low := strings.ToLower(str)
	r := strings.NewReplacer(
		"a", "1",
		"e", "2",
		"i", "3",
		"o", "4",
		"u", "5",
	)
	return r.Replace(low)
}

func Decode(str string) string {
	low := strings.ToLower(str)
	r := strings.NewReplacer(
		"1", "a",
		"2", "e",
		"3", "i",
		"4", "o",
		"5", "u",
	)
	return r.Replace(low)
}
