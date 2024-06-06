package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Exercise 002")

	// reading input from console
	var input int
	fmt.Print("Please enter a number : ")
	_, err := fmt.Scanln(&input)

	// checking for error and low 0
	if err != nil {
		log.Fatal("Please enter a number")
	}

	result, err := Ex002(input)

	if err != nil {
		log.Fatalf("Error for input %v: %v", input, err)
	}

	fmt.Printf("Factorial of %v = %v", input, result)
}

func Ex001(low, high int) string {
	var numbers []string
	for i := low; i <= high; i++ {
		if i%7 == 0 && i%5 != 0 {
			numbers = append(numbers, strconv.Itoa(i))
		}
	}
	return strings.Join(numbers, ",")
}

func Ex002(input int) (uint64, error) {
	if input < 0 {
		return 0, fmt.Errorf("Input must be a positive number")
	}
	var result uint64 = 1
	for i := 1; i <= input; i++ {
		result *= uint64(i)
	}
	return result, nil
}
