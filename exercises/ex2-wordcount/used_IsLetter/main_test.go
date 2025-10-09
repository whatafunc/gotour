package main

import (
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	testCases := []string{
		"I am learning Go!",
		"Hello world hello",
		"Multiple   spaces    here",
		"", // empty string
		"Punctuation! Should, work? Right.",
	}

	for _, tc := range testCases {
		fmt.Printf("Text: %q\n", tc)
		fmt.Printf("Count: %v\n\n", WordCount(tc))
	}
}
