package app

import (
	"fmt"
	"url-shortener/hasher"
	"url-shortener/shortener"
	"url-shortener/stores"

	"github.com/pkg/errors"
)

// URLShortener implements the Shortener interface
type URLShortener struct {
	store  stores.Store
	hasher hasher.Hasher
}

func (us *URLShortener) StoreURL(url shortener.URL) (shortener.ShortCode, error) {
	shortCode := us.hasher.Hash(url)
	err := us.store.StoreURL(shortCode, url)
	if err != nil {
		return shortener.ShortCode(""), errors.Wrap(err,
			fmt.Sprintln("Could not store url", url))
	}
	return shortCode, nil
}

func (us *URLShortener) GetURL(shortCode shortener.ShortCode) (shortener.URL, error) {
	return us.store.GetURL(shortCode)
}
