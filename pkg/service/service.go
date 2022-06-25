package service

import "github.com/khusainnov/weather/pkg/repository"

type Weather interface {
}

type Service struct {
	Weather
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
