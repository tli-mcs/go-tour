package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

const C = 50
const H = 30

func main() {
	fmt.Println(Ex010("hello world and practice makes perfect and hello world again"))
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

func Ex007(inputX int, inputY int) [][]int {
	result := make([][]int, inputX)
	for i := 0; i < inputX; i++ {
		result[i] = make([]int, inputY)
		for j := 0; j < inputY; j++ {
			result[i][j] = i * j
		}
	}
	return result
}

func Ex008(input string) string {
	arr := strings.Split(input, ",")

	sort.Strings(arr)
	return strings.Join(arr, ",")
}

func Ex009(input []string) []string {
	output := make([]string, len(input))
	for i, v := range input {
		output[i] = strings.ToUpper(v)
	}
	return output
}

func Ex010(input string) string {
	arr := strings.Split(input, " ")
	sort.Strings(arr)
	seen := make(map[string]int)
	result := []string{}
	for _, v := range arr {
		if _, ok := seen[v]; !ok {
			seen[v] = 1
			result = append(result, v)
		}
	}

	return strings.Join(result, " ")
}
