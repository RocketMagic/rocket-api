package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

func RedisCli() (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", GetConfig("redis.host"), GetConfig("redis.port"))
	password := fmt.Sprintf("%s", GetConfig("redis.password"))
	port, _ := GetConfig("redis.password").(int)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       port,
	})

	_, err := client.Ping().Result()
	log.Println(client.Ping())
	return client, err
}
