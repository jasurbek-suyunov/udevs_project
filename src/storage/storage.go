package storage

import (
	"context"
	"time"
)

type StorageI interface {
	User() UserI
}

type UserI interface {
	FindUserByID(id int) (string, error)
}

type CacheStorageI interface {
	Redis() RedisI
}

type RedisI interface {
	Set(ctx context.Context, key, value string, expTime time.Duration) error
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (value string, err error)
	Contains(ctx context.Context, key string) (bool, error)
}
