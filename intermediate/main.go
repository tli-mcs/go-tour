package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const C = 50
const H = 30

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(Ex006(text))
}

func Ex006(input string) map[int]int {
	arr := strings.Split(input, ",")
	result := make(map[int]int)
	for _, v := range arr {
		v = strings.Trim(v, " ")
		i, _ := strconv.Atoi(v)
		result[i] = int(math.Round(math.Sqrt(float64(2 * C * i / H))))
	}
	return result
}
