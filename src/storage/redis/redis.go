package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/jasurbek-suyunov/udevs_project/config"
	"github.com/jasurbek-suyunov/udevs_project/src/storage"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

type RedisCache struct {
	rdb *redis.Client

	cache storage.RedisI
}

// Redis implements storage.CacheStorageI
func (r *RedisCache) Redis() storage.RedisI {

	if r.cache == nil {
		r.cache = NewCache(r.rdb)
	}
	return r.cache
}

// consts for redis connection
const (
	readTimeout = 10 * time.Second // 10 seconds
)

	func NewRedisCache(cfg *config.Config) (storage.CacheStorageI, error) {

		// ...1: creating context
		var ctx context.Context = context.Background()

		val := os.Getenv("REDIS_POOL_SIZE")
		if val == "" {
			return nil, errors.New("REDIS_POOL_SIZE not set")
		}

		// ...2: opening connection to redis
		rdb := redis.NewClient(&redis.Options{
			Addr:        fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
			Password:    "",                      // no password set
			DB:          cast.ToInt(cfg.RedisDB), // use default DB
			PoolTimeout: readTimeout,
			PoolSize:    cast.ToInt(cfg.RedisPoolSize),
		})

		// ...3: checking connection
		pong := rdb.Ping(ctx)
		_, err := pong.Result()
		if err != nil {
			return nil, errors.New("cannot connect to redis")
		}

		// ...4: returning redis cache db
		return &RedisCache{
			rdb: rdb,
		}, nil
	}