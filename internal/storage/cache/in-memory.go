package cache

import (
	"link-shortener/internal/model"
	"link-shortener/internal/storage"
	"sync"
)

type InMemory struct {
	m sync.Map
}

func NewInMemory() storage.StorageInstance {
	return &InMemory{}
}

func (s *InMemory) Unique(shortUrl string, longUrl string) bool {
	v, ok := s.m.Load(shortUrl) //!ok-коротка ссылка ранее не использовалась,  (v.(string) == longUrl) - запись уже есть в хранилище
	return !ok || (v.(string) == longUrl)
}

// TODO:реализовать до конца выдачу уникального значения
func (s *InMemory) Save(longUrl string, shortUrl string) error {
	s.m.Store(shortUrl, longUrl)
	return nil
}

func (s *InMemory) Get(shortUrl string) (string, error) {
	v, ok := s.m.Load(shortUrl)
	if !ok {
		return "", model.ErrNotFound
	}

	longUrl := v.(string)

	return longUrl, nil
}
