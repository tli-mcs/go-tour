// Contexts - Canceling a context (withDeadline)
package main

import (
	"context"
	"fmt"
	"time"
)

// Modify the doSomething function
func doSomething(ctx context.Context) {
	// Deadline of 1,5 seconds
	deadline := time.Now().Add(1500 * time.Millisecond)
	// Repeat the same operation for cancelling the context, but this time instead of WithCancel() we will use WithDeadline(a,b) the a argument will be the parent context ctx
	// and the b argument will be the deadline above (1,5 seconds in this case)
	ctx, cancelCtx := context.WithDeadline(ctx, deadline)
	// defer cancelCtx when time has passed
	defer cancelCtx()
	// Make a new unbuffered channel of integers and assign it to printCh
	printCh := make(chan int)
	// call the doAnother function as goroutine
	go doAnother(ctx, printCh)
	// Modify the loop, for each number, wait for one second.
	// And whenever it receives a ctx.Done(), just break the selection.
outerLoop:
	for num := 1; num <= 5; num++ {
		select {
		// case 1 (receive a number to printCh channel)
		// then sleep for a second
		case printCh <- num:
			time.Sleep(500 * time.Millisecond)
		// case 2 (received ctx.done())
		case <-ctx.Done():
			fmt.Printf("doSomething: %s\n", ctx.Err())
			break outerLoop
		}
	}

	// sleep for 100 ms
	time.Sleep(1000 * time.Millisecond)
	// print that doSomething has finished
	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		// first case will be reciving a ctx.Done() call, if this happens we will handle errors and abort the doAnother function.
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		// Second case will receive a value from printCh and assign that to a variable called num, after that print the num variable
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

func main() {
	ctx := context.Background()
	doSomething(ctx)
}
