package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Exercise 003")
	var n int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal("Error occured: ", err)
	}

	fmt.Printf("%v", Ex003(n))
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

func Ex003(input int) map[int]int {
	result := make(map[int]int)
	for i := 1; i <= input; i++ {
		result[i] = i * i
	}
	return result
}
