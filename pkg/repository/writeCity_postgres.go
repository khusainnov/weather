package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type WritePostgres struct {
	db *sqlx.DB
}

func NewWritePostgres(db *sqlx.DB) *WritePostgres {
	return &WritePostgres{db: db}
}

func (wr *WritePostgres) WriteCity(city string) (int, error) {
	var id int
	if city == "" {
		return 0, errors.New("city line is empty")
	}
	query := fmt.Sprintf("INSERT INTO %s (city, counter) VALUES ($1, 1) RETURNING id;", searchQueries)
	req := wr.db.QueryRow(query, city)
	if err := req.Scan(&id); err != nil {
		query = fmt.Sprintf("UPDATE %s SET counter=counter + 1 WHERE city=$1;", searchQueries)
		_, err = wr.db.Query(query, city)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}
