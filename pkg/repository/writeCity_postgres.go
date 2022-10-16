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
	var exists bool
	if city == "" {
		return 0, errors.New("city line is empty")
	}

	queryExist := fmt.Sprintf("SELECT exists (SELECT 1 FROM %s WHERE city=$1);", searchQueries)
	queryCreate := fmt.Sprintf("INSERT INTO %s (city, counter) VALUES ($1, 1) RETURNING id;", searchQueries)
	queryUpdate := fmt.Sprintf("UPDATE %s SET counter=counter + 1 WHERE city=$1;", searchQueries)

	err := wr.db.QueryRow(queryExist, city).Scan(&exists)
	if err != nil {
		return 0, err
	}

	if exists == true {
		_, err = wr.db.Query(queryUpdate, city)
		if err != nil {
			return 0, err
		}
	} else {
		err = wr.db.QueryRow(queryCreate, city).Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}
