package storage

import "sync"

type Storage struct {
	mu   sync.Mutex
	data map[int]string
}

func (s *Storage) Store(jobID int, status string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[jobID] = status
}

func (s *Storage) Range(fn func(JobId int, status string) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for k, v := range s.data {
		fn(k, v)
	}
}

func New() *Storage {
	return &Storage{
		data: make(map[int]string),
	}
}
