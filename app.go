package main

import (
	"dfg/redis_demo"
	"fmt"
	"time"
)

func main() {
	redis_demo.Sub("test", func(msg string) {
		fmt.Println("receive " + msg)
	})

	redis_demo.Pub("test", "This is a test")

	time.Sleep(1 * time.Second)
}
