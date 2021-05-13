package main

import (
	"fmt"
)

var count int

func inc() {
	count += 1
}

func main() {
	for i := 0; i < 1000; i++ {
		inc()
	}
	fmt.Println(count)
}
