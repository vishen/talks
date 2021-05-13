package main

import "fmt"

func main() {
	// Create a new channel with 3 entries
	bufferedChan := make(chan int, 3)

	// Send 3 entries to the channel
	bufferedChan <- 1
	bufferedChan <- 2
	bufferedChan <- 3

	// If we try and send to the channel now when
	// the channel is full, we will block.
	// bufferedChan <- 4
	close(bufferedChan)

	for i := range bufferedChan {
		fmt.Println(i)
	}

}
