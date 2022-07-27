package service

import (
	"github.com/khusainnov/weather/pkg/repository"
)

type Authorization interface {
}

type Weather interface {
	WriteCity(city string) (int, error)
}

type Service struct {
	Authorization
	Weather
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Weather: NewWriteService(repos),
	}
}
