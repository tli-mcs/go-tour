// Exercise: Channels directions (only send/tx)

// Make a goroutine with a channel for only send data.
// The function should be called "send" and the send-only channel should be it's 1st and only argument
// Receive data from that channel is prohibited / will cause compiler errors

package main

import "fmt"

func send(c chan<- string) {
	c <- "Only send"
}

func main() {
	var c chan string = make(chan string, 2)
	c <- "hahah"
	send(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
