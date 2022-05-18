package runtime

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var defaultRuntime Runtime

type Runtime struct {
	db *gorm.DB
	rc *redis.Client
}

func (r *Runtime) GetDB() *gorm.DB {
	return nil
}

func (r *Runtime) GetRedisClient() *redis.Client {
	return nil
}

func SetDB(db *gorm.DB) {
	defaultRuntime.db = db
}

func SetRedisClient(rc *redis.Client) {
	defaultRuntime.rc = rc
}

func GetDB() *gorm.DB {
	return defaultRuntime.GetDB()
}

func GetRedisClient() *redis.Client {
	return defaultRuntime.GetRedisClient()
}
