package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	searchQueries = "search_queries"
)

type ConfigPG struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg ConfigPG) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
