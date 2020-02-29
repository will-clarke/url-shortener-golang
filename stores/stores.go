package stores

import (
	"url-shortener/shortener"
)

type Store interface {
	StoreURL(shortener.ShortCode, shortener.URL) error
	GetURL(shortener.ShortCode) (shortener.URL, error)
}

func NewInMemoryStore() InMemoryStore {
	s := InMemoryStore{}
	s.hashmap = make(map[shortener.ShortCode]shortener.URL)
	return s
}
