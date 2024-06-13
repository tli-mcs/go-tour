package main

import "fmt"

type Hotel struct {
	numRooms   int32
	streetName string
	hasPool    bool
}

func main() {
	h := Hotel{
		numRooms:   100,
		streetName: "Main St",
		hasPool:    true,
	}
	fmt.Println(h)
}
