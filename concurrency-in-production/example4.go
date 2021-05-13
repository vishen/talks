package main

import (
	"fmt"
	"sync"
)

// START OMIT
var count int

func inc(num int) {
	for i := 0; i < num; i++ {
		count += 1
	}
}

func main() {
	wg := sync.WaitGroup{}

	// Should match the number of goroutines being spawned
	wg.Add(1)

	go func() {
		defer wg.Done()
		inc(1000)
	}()

	fmt.Println(count)
	wg.Wait()
	fmt.Println(count)
}

// END OMIT
