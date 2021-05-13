package main

import "fmt"

func main() {
	var myInt32 int32
	var myString string
	var myBool bool
	var myByte byte

	fmt.Printf("myInt32=%v\n", myInt32)
	fmt.Printf("myString=%q\n", myString)
	fmt.Printf("myBool=%v\n", myBool)
	fmt.Printf("myByte=%v\n", myByte)

	myInt32 = 42
	myString = "this is my string"
	myBool = true
	myByte = 'a'

	fmt.Printf("myInt32=%v\n", myInt32)
	fmt.Printf("myString=%q\n", myString)
	fmt.Printf("myBool=%v\n", myBool)
	fmt.Printf("myByte=%v\n", myByte)
}
