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

func (cr *CheckRedis) CheckCacheCity(city string) (int64, error) {
	resp, err := cr.rdb.Exists(ctx, city).Result()
	if err != nil {
		return 0, err
	}

	return resp, nil
}

func (cr *CheckRedis) GetCacheCity(city string) ([]byte, error) {
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
