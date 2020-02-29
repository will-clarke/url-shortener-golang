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
	Store  stores.Store
	Hasher hasher.Hasher
}

func (us *URLShortener) StoreURL(url shortener.URL) (shortener.ShortCode, error) {
	shortCode := us.Hasher.Hash(url)
	validationErr := url.Validate()
	if validationErr != nil {
		return shortener.ShortCode(""), errors.Wrap(validationErr,
			fmt.Sprintln(string(url), "is not a valid URL"))

	}
	err := us.Store.StoreURL(shortCode, url)
	if err != nil {
		return shortener.ShortCode(""), errors.Wrap(err,
			fmt.Sprintln("Could not store url", url))
	}
	return shortCode, nil
}

func (us *URLShortener) GetURL(shortCode shortener.ShortCode) (shortener.URL, error) {
	return us.Store.GetURL(shortCode)
}
