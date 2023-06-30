package mono_redis

import (
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var (
	redisOnce   sync.Once
	redisClient *redis.Client
)

func ConnectRedis() *redis.Client {

	// go-redis包自带了连接池，会自动维护redis连接，因此创建一次client即可，不要查询一次redis就关闭client
	redisOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Addr:         "anner.wang:16379",
			Password:     "Anner_login_123",
			DB:           0,
			MaxRetries:   10,
			PoolSize:     50,
			MinIdleConns: 0,
			DialTimeout:  10 * time.Second,
		})
	})

	return redisClient
}
