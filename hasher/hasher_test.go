package hasher_test

import (
	"testing"
	"url-shortener/hasher"
	"url-shortener/shortener"
)

func TestHash(t *testing.T) {
	url := shortener.URL("https://example.com/foo")
	code := hasher.Hash(url)
	if len(code) != hasher.SIZE {
		t.Error("hahsed", url, "and got unexpected code", code)
	}
}
