package app

import "github.com/pkg/errors"

// URLShortener implements the Shortener interface
type URLShortener struct {
	store  store.Store
	hasher Hasher
}

func (us *URLShortener) StoreURL(url URL) (ShortCode, error) {
	shortCode := us.hasher.Hash(url)
	err := us.store.StoreURL(shortCode, url)
	if err != nil {
		return ShortCode{}, errors.Wrap(err, "Could not store url", url)
	}
	return shortCode, nil
}

func (us *URLShortener) GetURL(shortCode ShortCode) (URL, error) {
	return us.store.GetURL(shortCode)
}

func (url *URL) validate() error {
	_, err := url.ParseRequestURI(string(url))
	return err
}
