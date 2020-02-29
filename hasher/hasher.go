package hasher

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"url-shortener/shortener"
)

const SIZE = 12
const TOP_SECRET_SALT = "super-top-secret-probs-an-env-var"

func Hash(url shortener.URL) shortener.ShortCode {
	hashedBytes := sha256.Sum256([]byte(url + TOP_SECRET_SALT))
	base64String := base64.StdEncoding.EncodeToString(hashedBytes[:])
	sanitisedString := strings.ReplaceAll(base64String, "/", "-")
	return shortener.ShortCode(sanitisedString)[0:SIZE]
}
