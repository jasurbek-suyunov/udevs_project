package postgres

import (
	"context"
	"database/sql"
	"github.com/jasurbek-suyunov/udevs_project/models"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// UploadProfileImage implements storage.UserI.
func (u *userRepo) UploadProfileImage(ctx context.Context, userID string, image string) error {
	query := `UPDATE users SET profile_image_url = $1 WHERE id = $2`
	_, err := u.db.Exec(query, image, userID)
	if err != nil {
		log.Printf("Method: UploadProfileImage, Error: %v", err)
		return err
	}
	return nil
}

// Search implements storage.UserI.
func (u *userRepo) Search(ctx context.Context, query string) ([]models.SearchResult, error) {
	var results []models.SearchResult

	searchQuery := `
    	SELECT 'user' AS type, id, username, full_name, bio, profile_image_url, NULL AS content, NULL AS created_at 
		FROM users 
		WHERE username ILIKE '%' || $1 || '%' OR full_name ILIKE '%' || $1 || '%' 
		UNION 
		SELECT 'twit' AS type, t.id, u.username, NULL AS full_name, NULL AS bio, NULL AS profile_image_url, t.content, t.created_at 
		FROM twits t 
		JOIN users u ON u.id = t.user_id 
		WHERE t.content ILIKE '%' || $1 || '%';

	`

	err := u.db.SelectContext(ctx, &results, searchQuery, query)
	if err != nil {
		log.Printf("Method: Search, Error: %v", err)
		return nil, err
	}
	return results, nil
}

// GetFollowersList implements storage.UserI.
func (u *userRepo) GetFollowersList(ctx context.Context, userID int) ([]models.Follow, error) {
	var users []models.Follow

	query := `SELECT u.id, u.username, u.full_name, u.bio, u.email, u.profile_image_url, u.created_at
	FROM users u
	JOIN followers f ON u.id = f.follower_id
	WHERE f.following_id = $1`
	err := u.db.Select(&users, query, userID)
	if err != nil {
		log.Printf("Method: GetFollowersList, Error: %v", err)
		return nil, err
	}
	return users, nil
}

// GetFollowingList implements storage.UserI.
func (u *userRepo) GetFollowingList(ctx context.Context, userID int) ([]models.Follow, error) {
	var users []models.Follow

	query := `SELECT u.id, u.username, u.full_name, u.bio, u.email, u.profile_image_url, u.created_at
	FROM users u
	JOIN followers f ON u.id = f.following_id
	WHERE f.follower_id = $1`

	err := u.db.Select(&users, query, userID)
	if err != nil {
		log.Printf("Method: GetFollowingList, Error: %v", err)
		return nil, err
	}
	return users, nil
}

// FollowUser implements storage.UserI.
func (u *userRepo) FollowUser(ctx context.Context, followerID int, followingID int) error {
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
func (u *userRepo) GetFollowers(ctx context.Context, userID int) ([]models.Follow, error) {
	var users []models.Follow
	query := `SELECT u.id, u.username, u.full_name, u.bio, u.email , u.profile_image_url, u.created_at
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
func (u *userRepo) GetFollowing(ctx context.Context, userID int) ([]models.Follow, error) {
	var users []models.Follow
	query := `SELECT u.id, u.username, u.full_name, u.bio, u.email, u.profile_image_url, u.created_at
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

	resp := models.User{}

	query := `INSERT INTO users(username, full_name, bio, email, password_hash, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, username, full_name,bio, email, password_hash, created_at `

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

	if err != nil {
		log.Printf("Method: CreateUser, Error: %v", err)
		return nil, err
	}

	return &resp, nil
}

// DeleteUser implements storage.UserI
func (u *userRepo) DeleteUser(ctx context.Context, uresID string) error {

	query := `DELETE FROM users WHERE id = $1`

	result, err := u.db.ExecContext(ctx, query, uresID)

	if err != nil {
		log.Printf("Method: DeleteUser, Error: %v", err)
		return err
	}

	//check if user exists and affected count
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 || err != nil {
		return sql.ErrNoRows
	}

	return nil
}


// GetUserByUsername implements storage.UserI
func (u *userRepo) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {

	var result models.User

	query := `SELECT id, username, full_name,bio, email, profile_image_url, password_hash, created_at  FROM users WHERE username = $1`

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
		&result.ProfileImageURL,
		&result.PasswordHash,
		&result.CreatedAt,
	)

	if err != nil {
		log.Printf("Method: GetUserByUsername, Error: %v", err)
		return nil, err
	}

	return &result, nil
}

// UpdateUser implements storage.UserI
func (u *userRepo) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {

	var result models.User

	query := `UPDATE users SET username = $1, full_name = $2, bio = $3, email = $4 WHERE id = $5 RETURNING id, username, full_name,bio, email, password_hash, created_at `

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

	if err != nil {
		log.Printf("Method: UpdateUser, Error: %v", err)
		return nil, err
	}

	return &result, nil
}

// NewUserRepo creates a new user repository
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db}
}
