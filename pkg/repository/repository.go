package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Weather interface {
	WriteCity(city string) (int, error)
}

type Repository struct {
	Authorization
	Weather
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Weather: NewWritePostgres(db),
	}
}
