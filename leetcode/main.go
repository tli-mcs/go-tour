package main

import "strconv"

func twoSum(nums []int, target int) []int {
	for i, num := range nums {
		var rem = target - num
		for j, num2 := range nums {
			if num2 == rem && i != j {
				return []int{i, j}
			}
		}
		// You can access each index with i and each element with num
	}
	return []int{}
}

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	for i := 0; i < len(xStr)/2; i++ {
		if xStr[i] != xStr[len(xStr)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	twoSum([]int{2, 7, 11, 15}, 9)
}
