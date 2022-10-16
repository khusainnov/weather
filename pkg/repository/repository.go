package repository

import (
	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
}

type Weather interface {
	WriteCity(city string) (int, error)
}

type WeatherCache interface {
	WriteCacheCity(city string, wd []byte) error
	GetCacheCity(city string) ([]byte, error)
	CheckCacheCity(city string) (int64, error)
}

type Repository struct {
	Authorization
	Weather
	WeatherCache
}

func NewRepository(db *sqlx.DB, dbr *redis.Client) *Repository {
	return &Repository{
		Weather:      NewWritePostgres(db),
		WeatherCache: NewCheckRedis(dbr),
	}
}
