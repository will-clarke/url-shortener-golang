package shortener

import "net/url"

// Shortener is the where the main business logic lives

// URL is a URL that can be validated
type URL string

// ShortCode is a hash(?) of a URL
type ShortCode string

// Shortener is the interface external services
// (eg. CLI or Web Server) should use
type Shortener interface {
	StoreURL(URL) (ShortCode, error)
	GetURL(ShortCode) (URL, error)
}

func (u URL) Validate() error {
	_, err := url.ParseRequestURI(string(u))
	return err
}
