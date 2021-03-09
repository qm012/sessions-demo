package initialize

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

var rdbClient *redis.Client

func InitRedisClient() {

	rdbClient = redis.NewClient(&redis.Options{
		Addr:     "115.29.207.12:6379",
		Password: "123456",
		DB:       0,
		PoolSize: 1000,
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if _, err := rdbClient.Ping(ctx).Result(); err != nil {
		panic(err)
	}
	log.Print("redis 连接成功啦")
}
