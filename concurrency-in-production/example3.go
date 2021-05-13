package main

import (
	"fmt"
	"time"
)

// START OMIT
var count int

func inc(num int) {
	for i := 0; i < num; i++ {
		count += 1
	}
}

func main() {
	go inc(1000)
	fmt.Println(count)
	time.Sleep(5 * time.Second)
	fmt.Println(count)
}

// END OMIT
