package storage

import (
	"fmt"
	"net/http"
	"proj/internal/models"
	"sync"
)

type Storage struct {
	data []map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		data: make([]map[string]string, 0),
	}
}

func (s *Storage) Links(links models.LinksResponse) (map[string]string, int, error) {
	res := make(map[string]string)
	var wg sync.WaitGroup
	ch := make(chan map[string]string)
	for _, link := range links.Links {
		wg.Add(1)
		go checkLink(link, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		for k, v := range i {
			res[k] = v
		}
	}

	s.data = append(s.data, res)

	return s.data[len(s.data)-1], len(s.data), nil
}

func checkLink(link string, ch chan map[string]string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(link)
	if err != nil {
		ch <- map[string]string{
			link: "not available",
		}
		return
	}
	defer resp.Body.Close()
	fmt.Println(link, resp)

	if resp.StatusCode != 200 {
		ch <- map[string]string{
			link: "not available",
		}
	} else {
		ch <- map[string]string{
			link: "available",
		}
	}
}

func (s *Storage) GetAllLinks() ([]map[string]string, error) {
	return s.data, nil
}
