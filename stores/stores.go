package stores

import (
	"url-shortener/shortener"
)

type Store interface {
	StoreURL(shortener.ShortCode, shortener.URL) (shortener.ShortCode, error)
	GetURL(shortener.ShortCode) (shortener.URL, error)
}
