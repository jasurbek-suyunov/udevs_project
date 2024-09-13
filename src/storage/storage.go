package storage

import (
	"context"
	"jas/models"
	"time"
)

type StorageI interface {
	User() UserI
}

type UserI interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, urerID string) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
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
