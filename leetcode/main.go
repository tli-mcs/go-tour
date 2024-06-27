package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	// 去掉字符串首尾空格
	s = strings.TrimSpace(s)
	slice := strings.Split(s, " ")
	maxLen := 0
	for _, v := range slice {
		if len(v) > 0 {
			maxLen = len(v)
		}
	}
	return maxLen
}

func main() {
}
