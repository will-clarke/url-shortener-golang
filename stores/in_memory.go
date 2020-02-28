package stores

import "url-shortener/shortener"

type InMemoryStore struct {
	hashmap map[shortener.URL]shortener.URL
}

func (s InMemoryStore) StoreURL(shortCode shortener.ShortCode, url shortener.URL) (shortener.ShortCode, error) {
	return shortener.ShortCode("BLAH"), nil
}

func (s InMemoryStore) GetURL(shortCode shortener.ShortCode) (shortener.URL, error) {
	return shortener.URL("BLAH"), nil
}
