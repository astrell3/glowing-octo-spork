package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func AutoDetectAndConvert(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	if isMorseCode(input) {
		res := morse.ToText(input)
		return res, nil
	}
	res := morse.ToMorse(input)
	return res, nil
}

func isMorseCode(input string) bool {
	if len(input) == 0 {
		return false
	}
	words := strings.Split(input, "   ")

	for _, word := range words {
		chars := strings.Split(word, " ")

		for _, char := range chars {
			if char == "" {
				continue
			}

			if !isValidMorseChar(char) {
				return false
			}
		}
	}
	return len(input) > 0
}

func isValidMorseChar(char string) bool {
	for _, r := range char {
		if r != '.' && r != '-' {
			return false
		}
	}
	return true
}
