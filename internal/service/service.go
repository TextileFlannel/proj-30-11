package service

import (
	"proj/internal/models"
	"proj/internal/storage"
)

type Service struct {
	storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Links(links models.LinksResponse) (models.LinksRequest, error) {
	req, num, err := s.storage.Links(links)
	if err != nil {
		return models.LinksRequest{}, err
	}
	return models.LinksRequest{
		Links:    req,
		LinkNums: num,
	}, nil
}

func (s *Service) GetAllLinks() ([]map[string]string, error) {
	return s.storage.GetAllLinks()
}
