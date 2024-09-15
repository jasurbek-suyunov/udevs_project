package models

type Tweet struct {
	ID            string `json:"id" db:"id"`
	UserID        string `json:"user_id" db:"user_id"`
	Content       string `json:"content" db:"content"`
	MediaUrl      string `json:"media_url" db:"media_url"`
	LikesCount    int    `json:"likes_count" db:"likes_count"`
	RetweetsCount int    `json:"retweets_count" db:"retweets_count"`
	CreatedAt     int64  `json:"created_at" db:"created_at"`
}

type TweetRequest struct {
	Content string `json:"content" binding:"required" db:"content"`
}
