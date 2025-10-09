package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func(int) int {
	//sum := 0
	return func(x int) int {
		if x <= 1 {
			return x
		}
		r := Fibonacci()
		return r(x-1) + r(x-2)

	}
}

func main() {
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

/*
0
1
1
2
3
5
8
13
21
34

Program exited.

*/
