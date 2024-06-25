package main

import "fmt"

func counter(out chan<- int) {
	for i := 0; i < 11; i++ {
		out <- i
		fmt.Println("counter called, ch1 received val: ", i)
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {

	for i := range in {
		fmt.Println("squarer called, ch1 is sending val: ", i)
		fmt.Println("squarer called, ch2 is receiving val: ", i*i)
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println("ch2 sent val: ", i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
