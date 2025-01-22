package service

import (
	"strings"

	"github.com/izumiicxde/morse-code-decoder/utils"
)

func Generate(t string) string {
	var morseCode string
	for _, words := range strings.Split(strings.ToUpper(t), " ") {
		for _, char := range strings.Split(words, "") {
			if code, exists := utils.CharToMorseMap[char]; exists {
				morseCode = morseCode + code + " "
			}
		}
		morseCode += " /"
	}

	return morseCode[:len(morseCode)-1]
}
