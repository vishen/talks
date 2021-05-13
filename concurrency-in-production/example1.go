package main

import (
	"fmt"
)

var i int

func inc() {
	i += 1
}

func main() {
	for i := 0; i < 1000; i++ {
		inc()
	}
	fmt.Println(i)
}
