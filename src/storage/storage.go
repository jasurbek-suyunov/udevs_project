package storage

import (
	"context"
	"jas/models"
	"time"
)

type StorageI interface {
	User() UserI
	Tweet() TweetI
}

type UserI interface {
	// user
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, urerID string) error
	GetUserByID(ctx context.Context, id string) (*models.User, error)
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

type TweetI interface {
	// tweet
	CreateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error)
	UpdateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error)
	DeleteTweet(ctx context.Context, tweetID string) error
	GetTweetByID(ctx context.Context, id string) (*models.Tweet, error)
	GetTweetsByUserID(ctx context.Context, userID string) ([]models.Tweet, error)
	GetTweets(ctx context.Context, userID string) ([]models.Tweet, error)
	CanModifyTweet(ctx context.Context, userID string, tweetID string) (bool, error)
	LikeTweet(ctx context.Context, userID string, tweetID string) error
	UnlikeTweet(ctx context.Context, userID string, tweetID string) error
	RetweetTweet(ctx context.Context, userID string, tweetID string) error
	UnRetweetTweet(ctx context.Context, userID string, tweetID string) error
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