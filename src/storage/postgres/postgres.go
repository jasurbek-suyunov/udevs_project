package postgres

import (
	"fmt"
	"jas/config"
	"jas/src/storage"
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

func (s *Storage) User() storage.UserI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}
	return s.user
}

func (s *Storage) Tweet() storage.TweetI {
	if s.tweet == nil {
		s.tweet = NewTweetRepo(s.db)
	}
	return s.tweet
}
