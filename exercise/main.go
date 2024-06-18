// Exercise: Channels directions (only read/rx)

// Make a goroutine with a channel for only receive data.
// The function should be called "receive" and the receive-only channel should be it's 1st and only argument
// Sending data from that channel is prohibited / will cause compiler errors
// Feed some string into that channel.

package main

import "fmt"

func receive(c <-chan string) {
	fmt.Println(<-c)
}

func main() {
	var c chan string = make(chan string, 1)
	c <- "HelloWorld!"
	receive(c)
}
