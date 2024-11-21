package redis

import (
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(Addr string, Password string, Database int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       Database,
	})

	return rdb
}
