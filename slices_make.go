package main

import "fmt"

func main() {
	// Anonymous helper function to print slices.
	printSlice := func(l string, s []int) {
		fmt.Printf("%s: len=%d cap=%d %v\n", l, len(s), cap(s), s)
	}

	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	// You can still index into a slice, but only if it has a length.
	a[0] = 1
	printSlice("a", a)

	// Indexing into a slice with no length will cause a runtime panic.
	b[0] = 1
	printSlice("b", b)
}
