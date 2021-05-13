package main

import "fmt"

func main() {
	// Anonymous helper function to print slices.
	printSlice := func(i int, s []int) {
		fmt.Printf("%d: len=%d cap=%d %v\n", i, len(s), cap(s), s)
	}

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(1, s)
	printSlice(2, s[4:])
	printSlice(3, s[:4])

	// Nil slice is the default zero value for a slice.
	// A nil slice has a length and a capacity of 0, and
	// no underlying array.
	var emptySlice []int
	printSlice(4, emptySlice)
}
