package main

import "fmt"

func main() {
	word := "Hello, World"
	fmt.Println(len(word), word)
	// Indexing into a string gives us a uint8
	fmt.Printf("indexing: i=%d ch=%v type=%T\n", 2, word[2], word[2])

	// Iterating through a string using `for range`, each
	// element in the array is a int32.
	for i, ch := range word {
		fmt.Printf("for-range: i=%d ch=%v type=%T\n", i, ch, ch)
	}

	// This is to handle utf-8 encoding where a single encoded character
	// can be int32.
	word3 := "Hëllo"
	fmt.Println(len(word3), word3)
	fmt.Printf("indexing: i=%d ch=%v type=%T\n", 2, word3[2], word3[2])
	for i, ch := range word3 {
		fmt.Printf("for-range: i=%d ch=%v type=%T\n", i, ch, ch)
	}
}
