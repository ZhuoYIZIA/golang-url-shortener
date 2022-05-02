package database

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func RedisConnect() *redis.Client {
	Host := os.Getenv("REDIS_HOST")
	Port := os.Getenv("REDIS_PORT")
	Password := os.Getenv("REDIS_PASSWORD")

	return redis.NewClient(&redis.Options{
		Addr:     Host + ":" + Port,
		Password: Password, // no password set
		DB:       0,        // use default DB
	})
}
