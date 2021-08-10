package redisc

import (
	"log"

	"github.com/go-redis/redis/v7"
)

var DB *redis.Client

func init() {
	DB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "Redis12345679#", // no password set
		DB:       0,  // use default DB
	})

	pong, err := DB.Ping().Result()
	if err != nil {
		log.Printf("redisc init: %s", err)
	}
	log.Printf("redisc readied: %s", pong)
}


