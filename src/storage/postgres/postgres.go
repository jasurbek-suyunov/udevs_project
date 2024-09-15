package postgres

import (
	"fmt"
	"github.com/jasurbek-suyunov/udevs_project/config"
	"github.com/jasurbek-suyunov/udevs_project/src/storage"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cast"
)


type Storage struct {
	db *sqlx.DB
	user storage.UserI
	tweet storage.TweetI
}
// NewPostgres returns a new instance of the postgres storage.
func NewPostgres(cfg *config.Config) (storage.StorageI, error) {
	psqlConnString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
	)
	db, err := sqlx.Open("postgres", psqlConnString)
	if err != nil {
		log.Fatalf("cannot connect to postgresql db: %s", err.Error())
	}

	// setting some settings for db
	db.SetConnMaxLifetime(time.Duration(time.Duration(cast.ToInt(cfg.PostgresConnMaxIdleTime)).Minutes()))
	db.SetMaxOpenConns(cast.ToInt(cfg.PostgresMaxConnections))

	// checking for Ping&Pong
	if err := db.Ping(); err != nil {
		log.Fatalf("ping error %s", err.Error())
	}

	return &Storage{
		db: db,
	}, nil
}
// User returns a new instance of the user repository.
func (s *Storage) User() storage.UserI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}
	return s.user
}
// Tweet returns a new instance of the tweet repository.
func (s *Storage) Tweet() storage.TweetI {
	if s.tweet == nil {
		s.tweet = NewTweetRepo(s.db)
	}
	return s.tweet
}
