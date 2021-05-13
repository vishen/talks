package main

import (
	"fmt"
	"sync"
)

// START OMIT
var (
	count int
	mu    sync.Mutex // HL
)

func inc(num int) {
	mu.Lock()         // HL
	defer mu.Unlock() // HL
	for i := 0; i < num; i++ {
		count += 1
	}
}

// END OMIT
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		// We always want to Add(1) for every goroutine we start
		wg.Add(1)
		go func(i int) {
			// We always need to call Done() when the goroutine finishes
			defer wg.Done()
			inc(1000)
			fmt.Printf("finished: %d\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println(count)
}
