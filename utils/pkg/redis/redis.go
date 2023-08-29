package redis

import (
	"fmt"
	"rentbook/internal/config"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func newRedisClient(host string, password string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func InitRedis(cfg *config.AppConfig) *redis.Client {
	redisPort := strconv.Itoa(int(cfg.REDIS_PORT))
	var redisHost = fmt.Sprintf("%s:%s", cfg.REDIS_ADDRESS, redisPort)
	var redisPassword = cfg.REDIS_PASSWORD

	rdb := newRedisClient(redisHost, redisPassword)
	logrus.Info("REDIS client initialized")

	return rdb
}
