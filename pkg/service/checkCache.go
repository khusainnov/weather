package service

import (
	"github.com/khusainnov/weather/pkg/repository"
)

type CheckService struct {
	repo repository.WeatherCache
}

func NewCheckService(repo repository.WeatherCache) *CheckService {
	return &CheckService{repo: repo}
}

func (cs *CheckService) CheckCity(city string) ([]byte, error) {
	//var resp entity.Weather

	rb, err := cs.repo.CheckCity(city)
	if err != nil {
		return nil, err
	}

	/*err = json.Unmarshal(rb, &resp)
	if err != nil {
		return entity.Weather{}, err
	}*/

	return rb, nil
}

func (cs *CheckService) WriteCacheCity(city string, wb []byte) error {
	return cs.repo.WriteCacheCity(city, wb)
}
