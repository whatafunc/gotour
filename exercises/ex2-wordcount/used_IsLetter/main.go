package main

import (
	"strings"
	"unicode"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// Split using FieldsFunc for custom logic
	words := strings.FieldsFunc(s, func(r rune) bool {
		// Split on any non-letter character
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	counts := make(map[string]int)
	for _, word := range words {
		if word != "" {
			counts[strings.ToLower(word)]++ // Case insensitive
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
