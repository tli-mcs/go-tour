// Exercise: Pointers

// Print the address of a variable
package main

import "fmt"

func main() {
	var x int = 5
	// Your code goes here
	var p = &x
	fmt.Println(p)
}
