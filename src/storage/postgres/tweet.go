package postgres

import (
	"context"
	"fmt"
	"github.com/jasurbek-suyunov/udevs_project/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type tweetRepo struct {
	db *sqlx.DB
}

// RetweetTweet implements storage.TweetI.
func (t *tweetRepo) RetweetTweet(ctx context.Context, userID string, tweetID string) error {
	// Check if user has already retweeted this tweet
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM retweets WHERE user_id = $1 AND tweet_id = $2)`
	err := t.db.GetContext(ctx, &exists, query, userID, tweetID)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("User has already retweeted this tweet")
	}

	// If not retweeted, insert into retweets table and update retweets_count column
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO retweets (user_id, tweet_id, created_at) VALUES ($1, $2, $3)`
	_, err = tx.Exec(insertQuery, userID, tweetID, time.Now().Unix())
	if err != nil {
		tx.Rollback()
		return err
	}

	updateQuery := `UPDATE tweets SET retweets_count = retweets_count + 1 WHERE id = $1`
	_, err = tx.Exec(updateQuery, tweetID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// UnRetweetTweet implements storage.TweetI.
func (t *tweetRepo) UnRetweetTweet(ctx context.Context, userID string, tweetID string) error {
	// Check if user has retweeted this tweet
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM retweets WHERE user_id = $1 AND tweet_id = $2)`
	err := t.db.GetContext(ctx, &exists, query, userID, tweetID)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("User has not retweeted this tweet")
	}

	// If retweeted, delete from retweets table and update retweets_count column
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	deleteQuery := `DELETE FROM retweets WHERE user_id = $1 AND tweet_id = $2`
	_, err = tx.Exec(deleteQuery, userID, tweetID)
	if err != nil {
		tx.Rollback()
		return err
	}

	updateQuery := `UPDATE tweets SET retweets_count = retweets_count - 1 WHERE id = $1 AND retweets_count > 0`
	_, err = tx.Exec(updateQuery, tweetID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// LikeTweet implements storage.TweetI.
func (t *tweetRepo) LikeTweet(ctx context.Context, userID string, tweetID string) error {
	// Tekshiramiz, foydalanuvchi avval like qilganmi
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = $1 AND tweet_id = $2)`
	err := t.db.GetContext(ctx, &exists, query, userID, tweetID)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("User has already liked this tweet")
	}

	// Like qilmagan bo'lsa, likes jadvaliga yozib qo'yamiz va likes_count ustunini yangilaymiz
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO likes (user_id, tweet_id, created_at) VALUES ($1, $2, $3)`
	_, err = tx.Exec(insertQuery, userID, tweetID, time.Now().Unix())
	if err != nil {
		tx.Rollback()
		return err
	}

	updateQuery := `UPDATE tweets SET likes_count = likes_count + 1 WHERE id = $1`
	_, err = tx.Exec(updateQuery, tweetID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// UnlikeTweet implements storage.TweetI.
func (t *tweetRepo) UnlikeTweet(ctx context.Context, userID string, tweetID string) error {
	// Tekshiramiz, foydalanuvchi like qilganmi
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = $1 AND tweet_id = $2)`
	err := t.db.GetContext(ctx, &exists, query, userID, tweetID)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("User has not liked this tweet")
	}

	// Like qilgan bo'lsa, likes jadvalidan o'chirib, likes_count ni kamaytiramiz
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	deleteQuery := `DELETE FROM likes WHERE user_id = $1 AND tweet_id = $2`
	_, err = tx.Exec(deleteQuery, userID, tweetID)
	if err != nil {
		tx.Rollback()
		return err
	}

	updateQuery := `UPDATE tweets SET likes_count = likes_count - 1 WHERE id = $1 AND likes_count > 0`
	_, err = tx.Exec(updateQuery, tweetID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()

}
// CanModifyTweet implements storage.TweetI.
func (t *tweetRepo) CanModifyTweet(ctx context.Context, userID string, tweetID string) (bool, error) {
	// response
	resp := false

	// query
	query := `SELECT EXISTS(SELECT 1 FROM tweets WHERE id = $1 AND user_id = $2)`

	// exec and scan
	err := t.db.GetContext(ctx, &resp, query, tweetID, userID)
	if err != nil {
		return false, err
	}

	// return result if success
	return resp, nil
}

// GetTweetsByUserID implements storage.TweetI.
func (t *tweetRepo) GetTweetsByUserID(ctx context.Context, userID string) ([]models.Tweet, error) {
	// response
	resp := []models.Tweet{}

	// query
	query := `SELECT * FROM tweets WHERE user_id = $1`

	// exec and scan
	err := t.db.SelectContext(ctx, &resp, query, userID)
	if err != nil {
		return nil, err
	}

	// return result if success
	return resp, nil
}

// CreateTweet implements storage.TweetI.
func (t *tweetRepo) CreateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error) {
	// response
	resp := models.Tweet{}

	// query
	query := `INSERT INTO tweets (user_id,content , media_url, created_at) VALUES ($1, $2, $3, $4) RETURNING id,user_id,content,media_url,created_at`

	// exec and scan
	row := t.db.QueryRowContext(ctx, query, tweet.UserID, tweet.Content, tweet.MediaUrl, tweet.CreatedAt)
	err := row.Scan(
		&resp.ID,
		&resp.UserID,
		&resp.Content,
		&resp.MediaUrl,
		&resp.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	// return result if success
	return &resp, nil
}

// DeleteTweet implements storage.TweetI.
func (t *tweetRepo) DeleteTweet(ctx context.Context, tweetID string) error {
	// query
	query := `DELETE FROM tweets WHERE id = $1`

	// exec
	_, err := t.db.ExecContext(ctx, query, tweetID)
	if err != nil {
		return err
	}

	return nil
}

// GetTweetByID implements storage.TweetI.
func (t *tweetRepo) GetTweetByID(ctx context.Context, id string) (*models.Tweet, error) {
	// response
	resp := models.Tweet{}

	// query
	query := `SELECT * FROM tweets WHERE id = $1`

	// exec and scan
	err := t.db.GetContext(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}

	// return result if success
	return &resp, nil
}

// GetTweets implements storage.TweetI.
func (t *tweetRepo) GetTweets(ctx context.Context, userID string) ([]models.Tweet, error) {
	// response
	resp := []models.Tweet{}
	// query
	query := `SELECT * FROM tweets WHERE user_id = $1`

	// exec and scan
	err := t.db.SelectContext(ctx, &resp, query, userID)
	if err != nil {
		return nil, err
	}
	// return result if success
	return resp, nil
}

// UpdateTweet implements storage.TweetI.
func (t *tweetRepo) UpdateTweet(ctx context.Context, tweet *models.Tweet) (*models.Tweet, error) {

	// response
	resp := models.Tweet{}

	// query
	query := `UPDATE tweets SET content = $1, media_url = $2 WHERE id = $3 RETURNING id`

	// exec and scan
	row := t.db.QueryRowContext(ctx, query, tweet.Content, tweet.MediaUrl, tweet.ID)
	err := row.Scan(&resp.ID)
	if err != nil {
		return nil, err
	}

	// return result if success
	return &resp, nil

}
// NewTweetRepo creates a new instance of tweetRepo.
func NewTweetRepo(db *sqlx.DB) *tweetRepo {
	return &tweetRepo{
		db: db,
	}
}