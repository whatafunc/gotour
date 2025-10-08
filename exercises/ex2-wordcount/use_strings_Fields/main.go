package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	stat := make(map[string]int)
	for _, word := range words {
		stat[word]++ // case sensitive
	}

	return stat
}

func main() {
	wc.Test(WordCount)
}
