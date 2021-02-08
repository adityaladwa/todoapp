package cache

import "github.com/go-redis/redis/v8"

var redisClient *redis.Client

func Connect() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
