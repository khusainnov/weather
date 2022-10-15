package service

import (
	"github.com/khusainnov/weather/pkg/repository"
)

type Authorization interface {
}

type Weather interface {
	WriteCity(city string) (int, error)
}

type WeatherCache interface {
	WriteCacheCity(city string, wb []byte) error
	CheckCity(city string) ([]byte, error)
}

type Service struct {
	Authorization
	Weather
	WeatherCache
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Weather:      NewWriteService(repos),
		WeatherCache: NewCheckService(repos),
	}
}
