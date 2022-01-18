package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "101.33.127.167:6379",
		Password: "123456",
		DB:       0,
	})

	result, err := client.MSet("key1", "value1", "key2", "value2").Result()
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	} else {
		fmt.Printf("%v\n", result)
	}
	value1, _ := client.Get("key1").Result()
	fmt.Printf("%v\n", value1)
	value2, _ := client.Get("key2").Result()
	fmt.Printf("%v\n", value2)
}
