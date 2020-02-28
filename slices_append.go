package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)

	// append works on nil slices
	s = append(s, 1)
	printSlice(s)
	s = append(s, 2)
	printSlice(s)
	s = append(s, 3, 4, 5, 6)
	printSlice(s)
	s = append(s, 7)
	printSlice(s)

}
