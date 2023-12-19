package stores

import "git.sr.ht/~will-clarke/url-shortner-golang/shortener"

type InMemoryStore struct {
	hashmap map[shortener.ShortCode]shortener.URL
}

func (s *InMemoryStore) StoreURL(shortCode shortener.ShortCode, url shortener.URL) error {
	s.hashmap[shortCode] = url
	return nil
}

func (s *InMemoryStore) GetURL(shortCode shortener.ShortCode) (shortener.URL, error) {
	return s.hashmap[shortCode], nil
}
