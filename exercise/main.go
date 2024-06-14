// Exercise: Functions

// Create a variadic function that sums many numbers passed as arguments
// Variadic functions can be called with any number of arguments.

package main

import "fmt"

func calc(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	// Your code goes here
	fmt.Println(calc(1, 2, 3, 4, 5))
}
