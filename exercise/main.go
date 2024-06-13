package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Stat("go.mod")
	if err != nil {
		fmt.Printf("File does not exist\n")
	} else {
		fmt.Printf("File exists\n")
	}
}
