package main

import "fmt"

func main() {
	// Create a new unbuffered channel. This requires a goroutine
	// to be reading from the other end of the channel.
	unbufferedChan := make(chan int) // HL

	// Send some entries to the channel
	// TODO: Run in goroutine to fix
	unbufferedChan <- 1
	unbufferedChan <- 2
	unbufferedChan <- 3
	close(unbufferedChan)

	for i := range unbufferedChan {
		fmt.Println(i)
	}

}
