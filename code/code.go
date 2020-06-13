package code

import (
	"strings"
)

var encode = strings.NewReplacer(
	"a", "1",
	"e", "2",
	"i", "3",
	"o", "4",
	"u", "5",
)

var decode = strings.NewReplacer(
	"1", "a",
	"2", "e",
	"3", "i",
	"4", "o",
	"5", "u",
)

func Encode(str string) string {
	low := strings.ToLower(str)

	return encode.Replace(low)
}

func Decode(str string) string {
	low := strings.ToLower(str)

	return decode.Replace(low)
}
