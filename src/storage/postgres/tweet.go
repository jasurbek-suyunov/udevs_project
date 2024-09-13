package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
	"jas/models"
)

type tweetRepo struct {
	db *sqlx.DB
}

// CreateTweet implements storage.TweetI.
func (t *tweetRepo) CreateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error) {
	// check if image or video is present
	var media_url string 
	if tweet.ImageUrl != "" {
		media_url = tweet.ImageUrl
	} else {
		media_url = tweet.VideuUrl
	}
	// response 
	resp := models.Tweet{}

	// query
	query := `INSERT INTO tweets (user_id,content , media_url, created_at) VALUES ($1, $2, $3, $4) RETURNING id`

	// exec and scan
	row := t.db.QueryRowContext(ctx, query, tweet.UserID, tweet.Content, media_url, tweet.CreatedAt)
	err := row.Scan(&resp.ID)
	if err != nil {
		return nil, err
	}

	// return result if success
	return &resp, nil
}

// DeleteTweet implements storage.TweetI.
func (t *tweetRepo) DeleteTweet(ctx context.Context, tweetID string) error {
	panic("unimplemented")
}

// GetTweetByID implements storage.TweetI.
func (t *tweetRepo) GetTweetByID(ctx context.Context, id string) (*models.Tweet, error) {
	panic("unimplemented")
}

// GetTweets implements storage.TweetI.
func (t *tweetRepo) GetTweets(ctx context.Context, userID string) ([]models.Tweet, error) {
	panic("unimplemented")
}

// UpdateTweet implements storage.TweetI.
func (t *tweetRepo) UpdateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error) {
	panic("unimplemented")
}

func NewTweetRepo(db *sqlx.DB) *tweetRepo {
	return &tweetRepo{
		db: db,
	}
}
