package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

type CheckRedis struct {
	rdb *redis.Client
}

func NewCheckRedis(rdb *redis.Client) *CheckRedis {
	return &CheckRedis{rdb: rdb}
}

func (cr *CheckRedis) CheckCity(city string) ([]byte, error) {
	var rd []byte

	rd, err := cr.rdb.Get(ctx, city).Bytes()
	if err != nil {
		return nil, err
	}

	return rd, nil
}

func (cr *CheckRedis) WriteCacheCity(city string, wd []byte) error {
	err := cr.rdb.Set(ctx, city, wd, time.Second*60).Err()
	if err != nil {
		return err
	}

	return nil
}
