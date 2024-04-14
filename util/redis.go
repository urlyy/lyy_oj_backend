package util

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func RedisForgetPasswordKey(email string) string {
	return fmt.Sprintf("fp_%s", email)
}

func RedisChangeEmailKey(email string) string {
	return fmt.Sprintf("ce_%s", email)
}

func RedisTokenKey(userID int) string {
	return fmt.Sprintf("tk_%d", userID)
}

func RedisGet(key string) (string, error) {
	config := GetProjectConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       0,
	})
	val, err := client.Get(context.Background(), key).Result()
	return val, err
}
func RedisDel(key string) error {
	config := GetProjectConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       0,
	})
	err := client.Del(context.Background(), key).Err()
	return err
}

func RedisSet(key string, value interface{}, expireSec int) error {
	config := GetProjectConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       0,
	})
	err := client.Set(context.Background(), key, value, time.Second*time.Duration(expireSec)).Err()
	return err
}
