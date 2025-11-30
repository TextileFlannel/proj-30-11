package service

import (
	"bytes"
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

func (s *Service) Links(links models.LinksRequest) (models.LinksResponse, error) {
	return s.storage.CheckLinks(links)
}

func (s *Service) GetAllLinks() ([]map[string]string, error) {
	return s.storage.GetAllLinks()
}

func (s *Service) ReportLinks(linksList models.ReportLinksRequest) (bytes.Buffer, error) {
	return s.storage.ReportLinks(linksList)
}
