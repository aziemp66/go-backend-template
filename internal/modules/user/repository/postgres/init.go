package user_repository_postgres

import (
	user_repository "backend-template/internal/modules/user/repository"

	"github.com/jmoiron/sqlx"
)

type userRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) user_repository.UserRepository {
	return &userRepositoryPostgres{
		db: db,
	}
}
