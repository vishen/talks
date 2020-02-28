package main

import "fmt"

func main() {
	// Declaring an empty array.
	// Arrays have a fixed size in memory.
	var myArray [5]string

	myArray[1] = "Hello"
	myArray[3] = "World"

	fmt.Println(myArray)
	fmt.Printf("%#v\n", myArray)

	// Array Literal
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(primes)
}
