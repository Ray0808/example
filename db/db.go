package db

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func RedisConnect() {
	Client = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
	if err != nil {
		fmt.Println("Redis is not Connect")
	} else {
		fmt.Println("Redis is Connect")
	}
}
