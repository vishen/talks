package main

import "fmt"

func main() {
	myInt := 42
	myString := "this is my string"
	myBool := true
	myByte := 'a'

	fmt.Printf("myInt=%v type=%T\n", myInt, myInt)
	fmt.Printf("myString=%q type=%T\n", myString, myString)
	fmt.Printf("myBool=%v type=%T\n", myBool, myBool)
	fmt.Printf("myByte=%v type=%T\n", myByte, myByte)
}
