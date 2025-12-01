package storage

import (
	"sync"
)

type Storage struct {
	data   []map[string]string
	length int
	mu     sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		data:   make([]map[string]string, 0),
		length: 0,
	}
}

func (s *Storage) AddLink(link map[string]string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, link)
	s.length += 1
}

func (s *Storage) GetAllLinks() []map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.data
}

func (s *Storage) GetLength() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.length
}

func (s *Storage) GetByNums(nums []int) []map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]map[string]string, 0)
	for _, num := range nums {
		if num > 0 && num <= len(s.data) {
			result = append(result, s.data[num-1])
		} else {
			result = append(result, nil)
		}
	}
	return result
}
