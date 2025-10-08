package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ") // Only splits on spaces
	counts := make(map[string]int)

	for _, word := range words {
		if word != "" { // Check empty string on a space
			counts[word]++
		}
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
