package models

type Twit struct {
	ID            string `json:"id" db:"id"`
	UserID        string `json:"user_id" db:"user_id"`
	Content       string `json:"content" db:"content"`
	MediaUrl      string `json:"media_url" db:"media_url"`
	LikesCount    int    `json:"likes_count" db:"likes_count"`
	RetwitsCount int    `json:"retwits_count" db:"retwits_count"`
	CreatedAt     int64  `json:"created_at" db:"created_at"`
}

type TwitRequest struct {
	Content string `json:"content" binding:"required" db:"content"`
}
