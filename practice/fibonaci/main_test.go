package main

import (
	"testing"
)

func Test_main(t *testing.T) {
	f := Fibonacci()
	var results [10]int
	expected := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for i := 0; i < 10; i++ {
		results[i] = f(i)
		if results[i] != expected[i] {
			t.Errorf("Fibonacci sequence error at index %d: got %d, expected %d", i, results[i], expected[i])
		}
	}
}
