package dao

import (
	"github.com/go-redis/redis/v8"
	"github.com/zmicro-team/zmicro/core/config"
	"sync"
)

var (
	redisClient     *redis.Client
	onceRedisClient sync.Once
)

type Redis struct {
	Addr     string
	Password string
	DB       int
}

func GetRedisClient() *redis.Client {
	onceRedisClient.Do(func() {
		c := Redis{}
		config.Scan("redis", &c)
		redisClient = redis.NewClient(&redis.Options{
			Addr:     c.Addr,
			Password: c.Password,
			DB:       c.DB,
		})
	})

	return redisClient
}
