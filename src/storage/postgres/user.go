package postgres

import "github.com/jmoiron/sqlx"

type userRepo struct {
	db *sqlx.DB
}

// FindUserByID implements storage.UserI.
func (u *userRepo) FindUserByID(id int) (string, error) {
	panic("unimplemented")
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db}
}
