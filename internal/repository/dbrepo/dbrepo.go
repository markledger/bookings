package dbrepo

import (
	"database/sql"
	"github.com/markledger/bookings/internal/config"
	"github.com/markledger/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepository {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
