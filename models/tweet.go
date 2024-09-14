package models

type Tweet struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	Content   string `json:"content" db:"content"`
	MediaUrl string `json:"media_url" db:"media_url"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}

type TweetRequest struct {
	Content  string `json:"content" binding:"required" db:"content"`
	MediaUrl string `json:"media_url" db:"media_url"`
}