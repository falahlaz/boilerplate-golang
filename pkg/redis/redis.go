package redis

import (
	"fmt"
	"time"

	"github.com/falahlaz/boilerplate-golang/pkg/config"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func NewRedisClient() (*redis.Client, error) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.Config.Redis.Host, config.Config.Redis.Port),
		Password:     config.Config.Redis.Pass,
		DB:           0,
		PoolSize:     config.Config.Redis.PoolSize,
		MaxConnAge:   time.Duration(config.Config.Redis.MaxConnAge) * time.Second,
		MinIdleConns: config.Config.Redis.MinIdleConns,
	})

	return RedisClient, nil
}
