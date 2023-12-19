package hasher_test

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
	"testing"
	"git.sr.ht/~will-clarke/url-shortner-golang/hasher"
	"git.sr.ht/~will-clarke/url-shortner-golang/shortener"
)

const validChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+-"

func TestHashSHA256(t *testing.T) {
	hasherInstance := hasher.SHA256{}
	urls := []shortener.URL{
		"",
		"https://www.example.com/legit-url",
		"https://www.example.com/another-url",
		"any-other-random-strings",
		shortener.URL(randomString(1)),
		shortener.URL(randomString(10)),
		shortener.URL(randomString(100)),
	}
	for _, url := range urls {
		t.Run("TestHash-"+string(url), func(t *testing.T) {
			code := hasherInstance.Hash(url)
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

func randomString(l int) string {
	bytes := make([]byte, l)
	_, _ = rand.Read(bytes[:])
	return base64.StdEncoding.EncodeToString(bytes[:])
}
