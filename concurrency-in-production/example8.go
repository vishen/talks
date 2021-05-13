package main

import "fmt"

func main() {
	// Create a new channel with 3 entries
	bufferedChan := make(chan int, 3)

	// Write to the channel in the background
	go func() {
		for i := 0; i < 10; i++ {
			bufferedChan <- i
		}
		close(bufferedChan)
	}()

	for i := range bufferedChan {
		fmt.Println(i)
	}

}
