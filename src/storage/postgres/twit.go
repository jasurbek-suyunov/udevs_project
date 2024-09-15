package postgres

import (
	"context"
	"fmt"
	"github.com/jasurbek-suyunov/udevs_project/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type twitRepo struct {
	db *sqlx.DB
}

// RetwitTwit implements storage.TwitI.
func (t *twitRepo) RetwitTwit(ctx context.Context, userID string, twitID string) error {
	// Check if user has already retwited this twit
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM retwits WHERE user_id = $1 AND twit_id = $2)`
	err := t.db.GetContext(ctx, &exists, query, userID, twitID)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("User has already retwited this twit")
	}

	// If not retwited, insert into retwits table and update retwits_count column
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO retwits (user_id, twit_id, created_at) VALUES ($1, $2, $3)`
	_, err = tx.Exec(insertQuery, userID, twitID, time.Now().Unix())
	if err != nil {
		tx.Rollback()
		return err
	}

	updateQuery := `UPDATE twits SET retwits_count = retwits_count + 1 WHERE id = $1`
	_, err = tx.Exec(updateQuery, twitID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// UnRetwitTwit implements storage.TwitI.
func (t *twitRepo) UnRetwitTwit(ctx context.Context, userID string, twitID string) error {
	// Check if user has retwited this twit
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM retwits WHERE user_id = $1 AND twit_id = $2)`
	err := t.db.GetContext(ctx, &exists, query, userID, twitID)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("User has not retwited this twit")
	}

	// If retwited, delete from retwits table and update retwits_count column
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	deleteQuery := `DELETE FROM retwits WHERE user_id = $1 AND twit_id = $2`
	_, err = tx.Exec(deleteQuery, userID, twitID)
	if err != nil {
		tx.Rollback()
		return err
	}

	updateQuery := `UPDATE twits SET retwits_count = retwits_count - 1 WHERE id = $1 AND retwits_count > 0`
	_, err = tx.Exec(updateQuery, twitID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// LikeTwit implements storage.TwitI.
func (t *twitRepo) LikeTwit(ctx context.Context, userID string, twitID string) error {
	// Tekshiramiz, foydalanuvchi avval like qilganmi
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = $1 AND twit_id = $2)`
	err := t.db.GetContext(ctx, &exists, query, userID, twitID)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("User has already liked this twit")
	}

	// Like qilmagan bo'lsa, likes jadvaliga yozib qo'yamiz va likes_count ustunini yangilaymiz
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	insertQuery := `INSERT INTO likes (user_id, twit_id, created_at) VALUES ($1, $2, $3)`
	_, err = tx.Exec(insertQuery, userID, twitID, time.Now().Unix())
	if err != nil {
		tx.Rollback()
		return err
	}

	updateQuery := `UPDATE twits SET likes_count = likes_count + 1 WHERE id = $1`
	_, err = tx.Exec(updateQuery, twitID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// UnlikeTwit implements storage.TwitI.
func (t *twitRepo) UnlikeTwit(ctx context.Context, userID string, twitID string) error {
	// Tekshiramiz, foydalanuvchi like qilganmi
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM likes WHERE user_id = $1 AND twit_id = $2)`
	err := t.db.GetContext(ctx, &exists, query, userID, twitID)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("User has not liked this twit")
	}

	// Like qilgan bo'lsa, likes jadvalidan o'chirib, likes_count ni kamaytiramiz
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	deleteQuery := `DELETE FROM likes WHERE user_id = $1 AND twit_id = $2`
	_, err = tx.Exec(deleteQuery, userID, twitID)
	if err != nil {
		tx.Rollback()
		return err
	}

	updateQuery := `UPDATE twits SET likes_count = likes_count - 1 WHERE id = $1 AND likes_count > 0`
	_, err = tx.Exec(updateQuery, twitID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()

}
// CanModifyTwit implements storage.TwitI.
func (t *twitRepo) CanModifyTwit(ctx context.Context, userID string, twitID string) (bool, error) {
	// response
	resp := false

	// query
	query := `SELECT EXISTS(SELECT 1 FROM twits WHERE id = $1 AND user_id = $2)`

	// exec and scan
	err := t.db.GetContext(ctx, &resp, query, twitID, userID)
	if err != nil {
		return false, err
	}

	// return result if success
	return resp, nil
}

// GetTwitsByUserID implements storage.TwitI.
func (t *twitRepo) GetTwitsByUserID(ctx context.Context, userID string) ([]models.Twit, error) {
	// response
	resp := []models.Twit{}

	// query
	query := `SELECT * FROM twits WHERE user_id = $1`

	// exec and scan
	err := t.db.SelectContext(ctx, &resp, query, userID)
	if err != nil {
		return nil, err
	}

	// return result if success
	return resp, nil
}

// CreateTwit implements storage.TwitI.
func (t *twitRepo) CreateTwit(ctx context.Context, twit *models.Twit) (*models.Twit, error) {
	// response
	resp := models.Twit{}

	// query
	query := `INSERT INTO twits (user_id,content , media_url, created_at) VALUES ($1, $2, $3, $4) RETURNING id,user_id,content,media_url,created_at`

	// exec and scan
	row := t.db.QueryRowContext(ctx, query, twit.UserID, twit.Content, twit.MediaUrl, twit.CreatedAt)
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

// DeleteTwit implements storage.TwitI.
func (t *twitRepo) DeleteTwit(ctx context.Context, twitID string) error {
	// query
	query := `DELETE FROM twits WHERE id = $1`

	// exec
	_, err := t.db.ExecContext(ctx, query, twitID)
	if err != nil {
		return err
	}

	return nil
}

// GetTwitByID implements storage.TwitI.
func (t *twitRepo) GetTwitByID(ctx context.Context, id string) (*models.Twit, error) {
	// response
	resp := models.Twit{}

	// query
	query := `SELECT * FROM twits WHERE id = $1`

	// exec and scan
	err := t.db.GetContext(ctx, &resp, query, id)
	if err != nil {
		return nil, err
	}

	// return result if success
	return &resp, nil
}

// GetTwits implements storage.TwitI.
func (t *twitRepo) GetTwits(ctx context.Context, userID string) ([]models.Twit, error) {
	// response
	resp := []models.Twit{}
	// query
	query := `SELECT * FROM twits WHERE user_id = $1`

	// exec and scan
	err := t.db.SelectContext(ctx, &resp, query, userID)
	if err != nil {
		return nil, err
	}
	// return result if success
	return resp, nil
}

// UpdateTwit implements storage.TwitI.
func (t *twitRepo) UpdateTwit(ctx context.Context, twit *models.Twit) (*models.Twit, error) {

	// response
	resp := models.Twit{}

	// query
	query := `UPDATE twits SET content = $1, media_url = $2 WHERE id = $3 RETURNING id`

	// exec and scan
	row := t.db.QueryRowContext(ctx, query, twit.Content, twit.MediaUrl, twit.ID)
	err := row.Scan(&resp.ID)
	if err != nil {
		return nil, err
	}

	// return result if success
	return &resp, nil

}
// NewTwitRepo creates a new instance of twitRepo.
func NewTwitRepo(db *sqlx.DB) *twitRepo {
	return &twitRepo{
		db: db,
	}
}