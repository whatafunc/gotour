package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var arr [][]uint8

	for i := 0; i < dy; i++ {
		row := make([]uint8, dx) // Create a row with dx elements
		for x := 0; x < dx; x++ {
			// Fill each pixel with some interesting value

			// row[x] = uint8(x * i) // see output img if you can :)
			// G115: integer overflow conversion int -> uint8
			row[x] = uint8(2 + i) // Example 2: XOR pattern
			// row[x] = uint8((x + y) / 2) // Example 3: Average
		}
		arr = append(arr, row)
		// fmt.Print(arr)
	}

	// a slice of dx 8-bit unsigned integers
	// length dy
	return arr
}

func main() {
	pic.Show(Pic) // expects a function
}
