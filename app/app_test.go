package app_test

import (
	"testing"
	"url-shortener/app"
	"url-shortener/shortener"
)

func TestURLShortener_StoreURL(t *testing.T) {
	u := app.NewInMemoryURLShortener()
	url := shortener.URL("https://example.com/foo")
	shortCode, err := u.StoreURL(url)
	if err != nil {
		t.Error("error while storing URL:", err)
	}
	expectedShortCode := "Ti0-MV4cifgD"
	if string(shortCode) != expectedShortCode {
		t.Error("unexpected Shortcode", shortCode,
			"was expecting", expectedShortCode)
	}
}

func TestURLShortener_GetURL(t *testing.T) {
	u := app.NewInMemoryURLShortener()
	url := shortener.URL("https://example.com/foo")
	shortCode, _ := u.StoreURL(url)
	retrievedURL, err := u.GetURL(shortCode)
	if err != nil {
		t.Error("error while getting URL:", err)
	}
	if url != retrievedURL {
		t.Error("expected to retrieve", url, "not", retrievedURL)
	}
}
