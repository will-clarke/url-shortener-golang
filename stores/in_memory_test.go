package stores_test

import (
	"testing"
	"url-shortener/shortener"
	"url-shortener/stores"
)

func TestInMemoryStore(t *testing.T) {
	store := stores.InMemoryStore{}

	url := shortener.URL("https://example.com/omgwow")
	shortCode = shortener.ShortCode("a-key")

	shortCode, err := store.StoreURL(shortCode, url)
	if err != nil {
		t.Error("Cannot store URL", err)
	}

	fetchedURL, err := store.GetURL(shortCode)
	if err != nil {
		t.Error("Cannot retrieve URL", err)
	}

	if url != fetchedURL {
		t.Error("original URL", url, "doesn't match", "fetchedURL")
	}
}
