package runtime

import (
	"github.com/go-redis/redis/v8"
	zdb "github.com/zmicro-team/zim/pkg/database/db"
	zredis "github.com/zmicro-team/zim/pkg/database/redis"
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/log"
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

func (r *Runtime) SetDB(db *gorm.DB) {
	defaultRuntime.db = db
}

func (r *Runtime) SetRedisClient(rc *redis.Client) {
	defaultRuntime.rc = rc
}

func GetDB() *gorm.DB {
	return defaultRuntime.GetDB()
}

func GetRedisClient() *redis.Client {
	return defaultRuntime.GetRedisClient()
}

func Setup() {
	if config.Get("mysql") != nil {
		c := zdb.Config{}
		if err := config.Scan("mysql", &c); err != nil {
			log.Fatal(err)
		}
		db, err := zdb.Open(&c)
		if err != nil {
			log.Fatal(err)
		}
		defaultRuntime.SetDB(db)
	}

	if config.Get("redis") != nil {
		c := zredis.Config{}
		if err := config.Scan("redis", &c); err != nil {
			log.Fatal(err)
		}
		rc, err := zredis.NewClient(&c)
		if err != nil {
			log.Fatal(err)
		}
		defaultRuntime.SetRedisClient(rc)
	}
}
