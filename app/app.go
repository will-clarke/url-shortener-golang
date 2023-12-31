package app

import (
	"fmt"

	"git.sr.ht/~will-clarke/url-shortener-golang/hasher"
	"git.sr.ht/~will-clarke/url-shortener-golang/shortener"
	"git.sr.ht/~will-clarke/url-shortener-golang/stores"
)

// URLShortener implements the Shortener interface
type URLShortener struct {
	Store  stores.Store
	Hasher hasher.Hasher
}

func NewInMemoryURLShortener() URLShortener {
	hashInstance := hasher.SHA256{}
	return URLShortener{
		Store:  stores.NewInMemoryStore(),
		Hasher: &hashInstance,
	}
}

func NewFileStore() (URLShortener, error) {
	hashInstance := hasher.SHA256{}
	store, err := stores.NewFileStore()
	if err != nil {
		return URLShortener{}, err
	}
	return URLShortener{
		Store:  store,
		Hasher: &hashInstance,
	}, nil
}

func (us *URLShortener) StoreURL(url shortener.URL) (shortener.ShortCode, error) {
	shortCode := us.Hasher.Hash(url)
	validationErr := url.Validate()
	if validationErr != nil {
		return shortener.ShortCode(""), fmt.Errorf("%s is not a valid URL: %w", string(url), validationErr)

	}
	err := us.Store.StoreURL(shortCode, url)
	if err != nil {
		return shortener.ShortCode(""), fmt.Errorf("Could not store url %s: %w", string(url), err)
	}
	return shortCode, nil
}

func (us *URLShortener) GetURL(shortCode shortener.ShortCode) (shortener.URL, error) {
	return us.Store.GetURL(shortCode)
}
