package shortener

// Shortener is the where the main business logic lives

// URL is a URL that can be validated
type URL string

// ShortCode is a hash(?) of a URL
type ShortCode string

// Shortener is the interface is the main way external services
// eg. CLI or Web Server
type Shortener interface {
	StoreURL(URL) (ShortCode, error)
	GetURL(ShortCode) (URL, error)
}

// TODO: validate URL in method
