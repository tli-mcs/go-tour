// Exercise: Channels (buffered)

// Channels that accept a single message (<-) are synchronous. But go also can use async channels, or buffered channels
// Create a buffered string channel with a capacity of 4
// Send directly 4 different strings to that channel.
// Use the pop_message function 4 times to unbuffer the channel and see how it works

package main

import (
	"fmt"
	"time"
)

func pop_message(c chan string) {
	msg := <-c
	fmt.Println(msg)
}
func main() {
	// Your code goes here
	var c chan string = make(chan string, 4)
	c <- "My"
	c <- "Name"
	c <- "is"
	c <- "Enin"

	pop_message(c)
	pop_message(c)
	pop_message(c)
	pop_message(c)

	// this sleep is in order to not exit the program sooner than the routine lifetime :)
	time.Sleep(1 * time.Second)
}
