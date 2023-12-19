package stores_test

import (
	"testing"
	"git.sr.ht/~will-clarke/url-shortner-golang/shortener"
	"git.sr.ht/~will-clarke/url-shortner-golang/stores"
)

func TestInMemoryStore(t *testing.T) {
	store := stores.NewInMemoryStore()

	url := shortener.URL("https://example.com/omgwow")
	shortCode := shortener.ShortCode("omg-key")

	storeURL(t, store, shortCode, url)
	expectToRetrieveURL(t, store, shortCode, url)
}

func TestInMemoryStoreWithSeveralURLs(t *testing.T) {
	store := stores.NewInMemoryStore()

	url1 := shortener.URL("https://example.com/omgwow")
	shortCode1 := shortener.ShortCode("omg-key")
	url2 := shortener.URL("https://example.com/megalolz")
	shortCode2 := shortener.ShortCode("lolz-key")
	url3 := shortener.URL("https://example.com/something-else")
	shortCode3 := shortener.ShortCode("something-key")

	storeURL(t, store, shortCode1, url1)
	storeURL(t, store, shortCode2, url2)

	expectToRetrieveURL(t, store, shortCode1, url1)
	expectToRetrieveURL(t, store, shortCode2, url2)

	storeURL(t, store, shortCode3, url3)
	expectToRetrieveURL(t, store, shortCode3, url3)

	expectToRetrieveURL(t, store, shortCode1, url1)
}

func expectToRetrieveURL(t *testing.T, store stores.Store, shortCode shortener.ShortCode, expectedURL shortener.URL) {
	fetchedURL, err := store.GetURL(shortCode)
	if err != nil {
		t.Error("Cannot retrieve URL", err)
	}

	if expectedURL != fetchedURL {
		t.Error("expected URL", expectedURL, "doesn't match", fetchedURL)
	}
}

func storeURL(t *testing.T, store stores.Store, shortCode shortener.ShortCode, url shortener.URL) {
	err := store.StoreURL(shortCode, url)
	if err != nil {
		t.Error("Cannot store URL", err)
	}
}
