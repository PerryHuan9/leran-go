package redisClient

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

var ctx = context.Background()

func StartRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:3369",
		Password: "",
		DB:       0,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis connect failed:", err)
		return
	}
	fmt.Println("redis connected.")

}

func GetClient() *redis.Client {
	return redisClient
}
