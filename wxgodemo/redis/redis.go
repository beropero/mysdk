package redis

import (
	"github.com/go-redis/redis"
)

var (
	config = &redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,  
	}
	RedisClient = redis.NewClient(config)
)
