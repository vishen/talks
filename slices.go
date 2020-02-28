package main

import "fmt"

func main() {
	// A slice is dynamically-size and is a flexible view
	// into an array.
	primes := [5]int{2, 3, 5, 7, 11}
	var mySliceIntoPrimes []int = primes[1:4]
	fmt.Println(primes)
	fmt.Println(mySliceIntoPrimes)

	// A slice does not store any data, it just describes
	// a section of the underlying array. But changing the elements
	// of a slice modifies the underlying array.
	mySliceIntoPrimes[0] = 100
	fmt.Println(primes)

	// Slice literal will create an array and then gives you the
	// slice reference to that underlying array.
	// (I am unsure if you have actual access to the underlying array)
	mySliceOfBools := []bool{true, false, false, true}
	fmt.Println(mySliceOfBools)
}
