package main

import (
	"os"
)

func main() {
	file, err := os.Create("create-file.txt")

	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("Germany\nFrance\nUSA\nSpain\nUK\n")
}
