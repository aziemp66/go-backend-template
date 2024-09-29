package user_repository_postgres

import (
	"github.com/jmoiron/sqlx"
)

type userRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) *userRepositoryPostgres {
	return &userRepositoryPostgres{
		db: db,
	}
}
