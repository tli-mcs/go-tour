package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Person struct {
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}

func main() {
	ctx := context.Background()
	// Ensure that you have Redis running on your system
	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// Ensure that the connection is properly closed gracefully
	defer rdb.Close()
	rdb.Del(ctx, "FOO")
	result, err := rdb.Get(ctx, "FOO").Result()
	if err != nil {
		fmt.Println("FOO not found")
	} else {
		fmt.Printf("FOO has value %s\n", result)
	}
}
