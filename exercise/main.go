package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter an integer: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Failed to read integer:", err)
	} else {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}

}
