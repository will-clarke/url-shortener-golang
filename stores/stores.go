package stores

import (
	"os"

	"git.sr.ht/~will-clarke/url-shortener-golang/shortener"
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

func NewFileStore() (*FileStore, error) {
	f, err := os.OpenFile("example-filestore-urls.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &FileStore{f: f}, nil
}
