// Exercise: Channels - Closing

// Create a string channel "c" (make it a buffered channel)
// Add 2 different strings directly into that channel.
// Close the channel with the close() statement and read a quote from the channel, Can you read it?

package main

import "fmt"

func main() {
	c := make(chan string, 2)
	c <- "Hello"
	c <- "World"
	fmt.Println(<-c)
	close(c)
	c <- "abc"
	fmt.Println(<-c) // panic: send on closed channel
}
