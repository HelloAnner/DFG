package main

import (
	"dfg/mono_redis"
	"fmt"
	"time"
)

func main() {
	mono_redis.Sub("test", func(msg string) {
		fmt.Println("receive " + msg)
	})

	mono_redis.Pub("test", "This is a test")

	time.Sleep(1 * time.Second)
}
