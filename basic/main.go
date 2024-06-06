package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Ex001(2000, 3200))
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
