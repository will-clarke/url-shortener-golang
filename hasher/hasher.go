package hasher

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"git.sr.ht/~will-clarke/url-shortner-golang/shortener"
)

const SIZE = 12
const TOP_SECRET_SALT = "super-top-secret-probs-an-env-var"

type Hasher interface {
	Hash(shortener.URL) shortener.ShortCode
}

type SHA256 struct{}

func (h *SHA256) Hash(url shortener.URL) shortener.ShortCode {
	hashedBytes := sha256.Sum256([]byte(url + TOP_SECRET_SALT))
	base64String := base64.StdEncoding.EncodeToString(hashedBytes[:])
	sanitisedString := strings.ReplaceAll(base64String, "/", "-")
	return shortener.ShortCode(sanitisedString)[0:SIZE]
}
