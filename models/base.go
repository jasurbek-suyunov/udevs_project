package models

import (
	"database/sql"
)

type Error struct {
	Error string `json:"error"`
}

type Message struct {
	Message string `json:"message"`
}

type Meta struct {
	Total       int `json:"total"`
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	TotalPage   int `json:"total_page"`
}

type GetAllResponse struct {
	Data  interface{} `json:"data"`
	Meta  *Meta       `json:"meta"`
	Error string      `json:"error"`
	Code  int         `json:"code"`
}

type DefaultResponse struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Code  int         `json:"code"`
}

type Token struct {
	UserId    string `json:"user_id"`
	UserAgent string `json:"user_agent"`
}

type SearchResult struct {
	Type         string         `db:"type"` 
	ID           int            `db:"id"`
	Username     string         `db:"username"`
	Content      sql.NullString `db:"content,omitempty"` 
	FullName     sql.NullString `db:"full_name,omitempty"`
	Bio          sql.NullString `db:"bio,omitempty"` 
	ProfileImage sql.NullString `db:"profile_image_url,omitempty"`
	CreatedAt    sql.NullInt32 `db:"created_at,omitempty"`
}