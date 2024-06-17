// Exercise: Routines

// Document yourself and ask what a goroutine is.
// Make a go routine of the counter() function
// right after calling the go routine, in the next line. call the same routine with another different number

package main

import (
	"fmt"
	"strconv"
	"time"
)

func counter(x int, y string) {
	for i := 0; i < x; i++ {
		fmt.Println(y + " print: " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Your code goes here
	go counter(10, "goroutine 1")
	counter(10, "main routine")
	go counter(10, "goroutine 2")

	// this sleep is in order to not exit the program sooner than the routine lifetime :)
	time.Sleep(20 * time.Second)
	fmt.Println("All finished")
}
