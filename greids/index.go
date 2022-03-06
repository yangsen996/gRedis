package greids

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

var (
	redisClient     *redis.Client
	redisClientOnce sync.Once
)

func Redis() *redis.Client {
	redisClientOnce.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     "127.0.0.1:6379",
			Username: "",
			Password: "",
			DB:       0,

			MaxRetries:      0, // 重试策略
			MinRetryBackoff: 8 * time.Millisecond,
			MaxRetryBackoff: 512 * time.Millisecond,

			DialTimeout:  5 * time.Second, //建立连接超时时间
			ReadTimeout:  3 * time.Second, //读超时
			WriteTimeout: 3 * time.Second, //写超时

			PoolSize:           15,
			MinIdleConns:       10, //最小链接数
			MaxConnAge:         0,
			PoolTimeout:        4 * time.Second, //等待空闲连接的最大等待时间
			IdleTimeout:        5 * time.Second, //闲置连接超时
			IdleCheckFrequency: 60 * time.Second,
		})
		pong, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalln("connect err", err)
		}
		log.Println(pong)
	})
	return redisClient
}
