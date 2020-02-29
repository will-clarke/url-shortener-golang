package hasher_test

import (
	"strings"
	"testing"
	"url-shortener/hasher"
	"url-shortener/shortener"
)

const validChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+-"

func TestHash(t *testing.T) {
	urls := []shortener.URL{
		"",
		"https://www.example.com/legit-url",
		"https://www.example.com/another-url",
		"any other string",
	}
	for _, url := range urls {
		t.Run("TestHash-"+string(url), func(t *testing.T) {
			code := hasher.Hash(url)
			if len(code) != hasher.SIZE {
				t.Error("hahsed", url, "and got unexpected code", code)
			}
			for _, c := range code {
				if !strings.Contains(validChars, string(c)) {
					t.Error("was expecting", string(c), "in", code,
						"to only contain these characters:", validChars)
				}

			}
		})
	}
}
