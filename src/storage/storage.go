package storage

import (
	"context"
	"github.com/jasurbek-suyunov/udevs_project/models"
	"time"
)

type StorageI interface {
	User() UserI
	Twit() TwitI
}

type UserI interface {
	// user
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	UploadProfileImage(ctx context.Context, userID string, image string) error
	DeleteUser(ctx context.Context, urerID string) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	// follow
	FollowUser(ctx context.Context, followerID, followingID int) error
	UnFollowUser(ctx context.Context, followerID, followingID int) error
	IsFollowing(ctx context.Context, followerID, followingID int) (bool, error)
	GetFollowers(ctx context.Context, userID int) ([]models.Follow, error)
	GetFollowing(ctx context.Context, userID int) ([]models.Follow, error)
	GetFollowingList(ctx context.Context, userID int) ([]models.Follow, error)
	GetFollowersList(ctx context.Context, userID int) ([]models.Follow, error)
	// search
	Search(ctx context.Context, query string) ([]models.SearchResult, error)
}

type TwitI interface {
	// twit
	CreateTwit(ctx context.Context, twit *models.Twit) (*models.Twit, error)
	UpdateTwit(ctx context.Context, twit *models.Twit) (*models.Twit, error)
	DeleteTwit(ctx context.Context, twitID string) error
	GetTwitByID(ctx context.Context, id string) (*models.Twit, error)
	GetTwitsByUserID(ctx context.Context, userID string) ([]models.Twit, error)
	GetTwits(ctx context.Context, userID string) ([]models.Twit, error)
	CanModifyTwit(ctx context.Context, userID string, twitID string) (bool, error)
	LikeTwit(ctx context.Context, userID string, twitID string) error
	UnlikeTwit(ctx context.Context, userID string, twitID string) error
	RetwitTwit(ctx context.Context, userID string, twitID string) error
	UnRetwitTwit(ctx context.Context, userID string, twitID string) error
}
type CacheStorageI interface {
	Redis() RedisI
}

type RedisI interface {
	// key value
	Set(ctx context.Context, key, value string, expTime time.Duration) error
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string) (value string, err error)
	Contains(ctx context.Context, key string) (bool, error)
}