package repository

import (
	"context"

	"github.com/go-redis/redis/v9"
)

type ConfigRedis struct {
	Port     string
	Password string
	DB       int
}

func NewRedisDB(cfg ConfigRedis, ctx context.Context) (*redis.Client, error) {
	dbr := redis.NewClient(&redis.Options{
		Addr:     cfg.Port,
		Password: cfg.Password,
		DB:       0,
	})

	_, err := dbr.Ping(ctx).Result()

	return dbr, err
}
