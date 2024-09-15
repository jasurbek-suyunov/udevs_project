package models

type User struct {
	ID           string `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	FullName     string `json:"full_name" db:"full_name"`
	Bio 		string `json:"bio" db:"bio"`
	Email        string `json:"email" db:"email"`
	ProfileImageURL string `json:"profile_image_url" db:"profile_image_url"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	CreatedAt    int64  `json:"created_at" db:"created_at"`
}

type UserSignUpRequest struct {
	Username        string `json:"username" binding:"required" db:"username"`
	FullName		string `json:"full_name" binding:"required" db:"full_name"`
	Bio 			string `json:"bio" binding:"required" db:"bio"`
	Email           string `json:"email" binding:"required" db:"email"`
	Password        string `json:"password" binding:"required" db:"password"`
	ConfirmPassword string `json:"confirm_password" binding:"required" db:"confirm_password"`
}

type UserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	Bio 	  string `json:"bio"`
	Email     string `json:"email"`
	ProfileImageURL string `json:"profile_image_url"`
	CreatedAt int64  `json:"created_at"`
}

type LoginResponse struct {
	Data  *UserResponse `json:"data"`
	Error string        `json:"error"`
	Code  int           `json:"code"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required" db:"username"`
	Password string `json:"password" binding:"required" db:"password"`
}

type FollowRequest struct {
	FollowedID int `json:"followed_id" binding:"required" db:"followed_id"`
}
type FollowResponse struct {
	FollowedID int `json:"followed_id" db:"followed_id"`
}

type UnFollowRequest struct {
	FollowedID int `json:"followed_id" binding:"required" db:"followed_id"`
}

type Follow struct {
	ID           string `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	FullName     string `json:"full_name" db:"full_name"`
	Bio 		string `json:"bio" db:"bio"`
	Email        string `json:"email" db:"email"`
	ProfileImageURL string `json:"profile_image_url" db:"profile_image_url"`
	CreatedAt    int64  `json:"created_at" db:"created_at"`
}