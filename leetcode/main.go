package main

import "sort"

type MinStack struct {
	stack       []int
	sortedStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.sortedStack = append(this.sortedStack, val)
	sort.Ints(this.sortedStack)
}

func (this *MinStack) Pop() {
	val := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	idx := -1
	for i, v := range this.sortedStack {
		if v == val {
			idx = i
			break
		}
	}
	this.sortedStack = append(this.sortedStack[:idx], this.sortedStack[idx+1:]...)
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.sortedStack[0]
}

func main() {
}
