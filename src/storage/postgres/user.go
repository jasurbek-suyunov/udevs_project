package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"jas/models"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

const (
	userTable  = "users"
	userFields = `id, username, full_name,bio, email, password_hash, created_at`
)


// FollowUser implements storage.UserI.
func (u *userRepo) FollowUser(ctx context.Context, followerID int, followingID int) error {
	fmt.Println("(((((((((((FollowerID))))))))))))))): ", followerID)
	query := `INSERT INTO followers (follower_id, following_id, created_at) 
	VALUES ($1, $2, $3)`
	_, err := u.db.Exec(query, followerID, followingID, time.Now().Unix())
	if err != nil {
		log.Printf("Method: FollowUser, Error: %v", err)
		return err
	}
	return nil
}

// GetFollowers implements storage.UserI.
func (u *userRepo) GetFollowers(ctx context.Context, userID int) ([]models.User, error) {
	var users []models.User
	query := `SELECT u.id, u.username, u.full_name, u.bio, u.email, u.password_hash, u.created_at
	FROM users u
	JOIN followers f ON u.id = f.follower_id
	WHERE f.following_id = $1`
	err := u.db.Select(&users, query, userID)
	if err != nil {
		log.Printf("Method: GetFollowers, Error: %v", err)
		return nil, err
	}
	return users, nil
}

// GetFollowing implements storage.UserI.
func (u *userRepo) GetFollowing(ctx context.Context, userID int) ([]models.User, error) {
	var users []models.User
	query := `SELECT u.id, u.username, u.full_name, u.bio, u.email, u.password_hash, u.created_at
	FROM users u
	JOIN followers f ON u.id = f.following_id
	WHERE f.follower_id = $1`
	err := u.db.Select(&users, query, userID)
	if err != nil {
		log.Printf("Method: GetFollowing, Error: %v", err)
		return nil, err
	}
	return users, nil
}

// IsFollowing implements storage.UserI.
func (u *userRepo) IsFollowing(ctx context.Context, followerID int, followingID int) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM followers WHERE follower_id = $1 AND following_id = $2)`

	err := u.db.QueryRow(query, followerID, followingID).Scan(&exists)
	if err != nil {
		log.Printf("Method: IsFollowing, Error: %v", err)
		return false, err
	}
	return exists, nil
}

// UnFollowUser implements storage.UserI.
func (u *userRepo) UnFollowUser(ctx context.Context, followerID int, followingID int) error {
	query := `DELETE FROM followers WHERE follower_id = $1 AND following_id = $2`
	_, err := u.db.Exec(query, followerID, followingID)
	if err != nil {
		log.Printf("Method: UnFollowUser, Error: %v", err)
		return err
	}
	return nil
}

// CreateUser implements repository.UserI
func (u *userRepo) CreateUser(ctx context.Context, usr *models.User) (*models.User, error) {

	// response model
	resp := models.User{}

	// query
	query := `INSERT INTO users(username, full_name,bio, email, password_hash, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING ` + userFields

	// exec and scan
	err := u.db.QueryRow(
		query,
		usr.Username,
		usr.FullName,
		usr.Bio,
		usr.Email,
		usr.PasswordHash,
		usr.CreatedAt,
	).Scan(
		&resp.ID,
		&resp.Username,
		&resp.FullName,
		&resp.Bio,
		&resp.Email,
		&resp.PasswordHash,
		&resp.CreatedAt,
	)

	// check if user exists
	if err != nil {
		log.Printf("Method: CreateUser, Error: %v", err)
		return nil, err
	}

	// return result if success
	return &resp, nil
}

// DeleteUser implements storage.UserI
func (u *userRepo) DeleteUser(ctx context.Context, uresID string) error {

	//query
	query := `DELETE FROM users WHERE id = $1`

	// exec
	result, err := u.db.ExecContext(ctx, query, uresID)

	// check if error
	if err != nil {
		log.Printf("Method: DeleteUser, Error: %v", err)
		return err
	}

	//check if user exists and affected count
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		return sql.ErrNoRows
	}

	// if success
	return nil
}

// GetUserByID implements storage.UserI
func (u *userRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {

	// response model
	var result models.User

	// query
	query := `SELECT ` + userFields + ` FROM users WHERE id = $1`

	// exec and scan
	if err := u.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&result.ID,
		&result.Username,
		&result.FullName,
		&result.Bio,
		&result.Email,
		&result.PasswordHash,
		&result.CreatedAt,
	); err != nil {
		log.Printf("Method: GetUserByID, Error: %v", err)
		return nil, err
	}

	// return result
	return &result, nil
}

// GetUserByUsername implements storage.UserI
func (u *userRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {

	// response model
	var result models.User

	// query
	query := `SELECT ` + userFields + ` FROM users WHERE username = $1`

	// exec and scan
	err := u.db.QueryRowContext(
		ctx,
		query,
		username,
	).Scan(
		&result.ID,
		&result.Username,
		&result.FullName,
		&result.Bio,
		&result.Email,
		&result.PasswordHash,
		&result.CreatedAt,
	)

	// check error
	if err != nil {
		log.Printf("Method: GetUserByUsername, Error: %v", err)
		return nil, err
	}

	// return result
	return &result, nil
}

// UpdateUser implements storage.UserI
func (u *userRepo) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {

	// response model
	var result models.User

	// query
	query := `UPDATE users SET username = $1, full_name = $2, bio = $3, email = $4 WHERE id = $5 RETURNING ` + userFields

	// exec and scan
	err := u.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.FullName,
		user.Bio,
		user.Email,
		user.ID,
	).Scan(
		&result.ID,
		&result.Username,
		&result.FullName,
		&result.Bio,
		&result.Email,
		&result.PasswordHash,
		&result.CreatedAt,
	)

	// check error
	if err != nil {
		log.Printf("Method: UpdateUser, Error: %v", err)
		return nil, err
	}

	// return result
	return &result, nil
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db}
}