package service

import (
	"github.com/khusainnov/weather/pkg/repository"
)

type WriteService struct {
	repo repository.Weather
}

func NewWriteService(repo repository.Weather) *WriteService {
	return &WriteService{repo: repo}
}

func (w *WriteService) WriteCity(city string) (int, error) {
	return w.repo.WriteCity(city)
}
