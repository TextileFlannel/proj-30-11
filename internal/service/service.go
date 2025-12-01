package service

import (
	"bytes"
	"proj/internal/checker"
	"proj/internal/models"
	"proj/internal/report"
	"proj/internal/storage"
	"sync"
)

type Service struct {
	storage   *storage.Storage
	checker   *checker.Checker
	generator *report.Generator
}

func NewService(storage *storage.Storage) *Service {
	return &Service{
		storage:   storage,
		checker:   checker.NewChecker(),
		generator: report.NewGenerator(),
	}
}

func (s *Service) Links(links models.LinksRequest) (models.LinksResponse, error) {
	res := make(map[string]string)
	var wg sync.WaitGroup
	ch := make(chan models.Result)

	for _, link := range links.Links {
		wg.Add(1)
		go func(l string) {
			result := s.checker.Check(l)
			ch <- result
			wg.Done()
		}(link)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		res[result.URL] = result.Status
	}

	s.storage.AddLink(res)

	return models.LinksResponse{
		Links:    res,
		LinkNums: s.storage.GetLength(),
	}, nil
}

func (s *Service) GetAllLinks() ([]map[string]string, error) {
	return s.storage.GetAllLinks(), nil
}

func (s *Service) ReportLinks(req models.ReportLinksRequest) (bytes.Buffer, error) {
	links := make([]models.LinksResponse, 0)
	data := s.storage.GetByNums(req.LinksList)

	for i, idx := range req.LinksList {
		if idx > 0 && idx <= s.storage.GetLength() && data[i] != nil {
			links = append(links, models.LinksResponse{
				LinkNums: idx,
				Links:    data[i],
			})
		}
	}

	return s.generator.Generate(links)
}
