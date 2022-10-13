package utils

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

var RDB = InitRedisDB()
var CTX = context.Background()

func InitRedisDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
func RedisSet(key, value string, duration time.Duration) error {
	err := RDB.Set(CTX, key, value, duration).Err()
	if err != nil {
		return err
	}
	return nil
}
func RedisGet(key string) (string, error) {
	val, err := RDB.Get(CTX, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
