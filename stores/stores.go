package stores

import (
	"git.sr.ht/~will-clarke/url-shortner-golang/shortener"
)

// Store interface is how the shortener interacts
// with different persistance mechanisms
type Store interface {
	StoreURL(shortener.ShortCode, shortener.URL) error
	GetURL(shortener.ShortCode) (shortener.URL, error)
}

func NewInMemoryStore() *InMemoryStore {
	s := InMemoryStore{}
	s.hashmap = make(map[shortener.ShortCode]shortener.URL)
	return &s
}
