package service

import (
	"strings"

	"github.com/izumiicxde/morse-code-decoder/utils"
)

// m is the morse code
func Decode(m string) string {
	var decodedString string
	if len(m) == 0 {
		return ""
	}

	words := strings.Split(m, "/")
	for _, word := range words {
		for _, letter := range strings.Split(word, " ") {
			decodedString += utils.MorseCodeMap[letter]
		}
		decodedString += " "
	}
	return decodedString
}
