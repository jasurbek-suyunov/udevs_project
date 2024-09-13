package models

type Tweet struct {
	ID        string `json:"id" db:"id"`
	UserID    string `json:"user_id" db:"user_id"`
	Content   string `json:"content" db:"content"`
	ImageUrl string `json:"image_url" db:"image_url"`
	VideuUrl string `json:"video_url" db:"video_url"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}

type TweetRequest struct {
	Content  string `json:"content" binding:"required" db:"content"`
	ImageUrl string `json:"image_url" db:"image_url"`
	VideoUrl string `json:"video_url" db:"video_url"`
}