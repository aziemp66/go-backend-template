package util_db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(host, user, password, dbName, port string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, port))
	if err != nil {
		panic(err)
	}

	return db
}

type PostgresDsn struct {
	Host     string
	Port     int
	User     string
	Password string
	Db       string
}

func (p PostgresDsn) ToString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", p.Host, p.User, p.Password, p.Db, p.Port)
}
