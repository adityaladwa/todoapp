package cache

import "github.com/go-redis/redis/v8"

var RedisClient *redis.Client

func Connect() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
