package redisclient

import "github.com/go-redis/redis"

func NewRedisCli(host, password string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
}
