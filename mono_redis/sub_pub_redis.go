package mono_redis

import (
	"fmt"
	redis "github.com/go-redis/redis"
)

func Sub(channel string, doSomething func(msg string)) {
	sub := ConnectRedis().Subscribe(channel)
	handleMsg(sub, doSomething)
}

func Pub(channel, msg string) {
	ConnectRedis().Publish(channel, msg)
}

// 方式一获取channel的消息
func handleMsg(sub *redis.PubSub, doSomething func(msg string)) {
	iface, err := sub.Receive()

	if err != nil {
		fmt.Println("build sub object fail")
	}

	switch iface.(type) {
	case *redis.Subscription:
		fmt.Println("sub success")
	case *redis.Message:
		doSomething(iface.(redis.Message).Payload)
	case *redis.Pong:
		fmt.Println("receive pong message")
	default:
		fmt.Println("something error")
	}
}
