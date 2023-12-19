package stores

import (
	"fmt"
	"io"
	"os"
	"strings"

	"git.sr.ht/~will-clarke/url-shortener-golang/shortener"
)

// Filestore stores URLs in a plain text file.
// This is super inefficient and useless.... but it is persistent
// which is good for demonstration
type FileStore struct {
	f *os.File
}

const absolutelyTerribleDelimiter = "==@=="

func (s *FileStore) StoreURL(shortCode shortener.ShortCode, url shortener.URL) error {
	// Ugh. Don't look too closely at this 🙈
	_, err := s.f.WriteString(string(shortCode) + absolutelyTerribleDelimiter + string(url) + "\n")
	return err
}

func (s *FileStore) GetURL(shortCode shortener.ShortCode) (shortener.URL, error) {
	file, err := io.ReadAll(s.f)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if strings.Contains(line, string(shortCode)) {
			splitString := strings.Split(line, absolutelyTerribleDelimiter)
			if len(splitString) != 2 {
				return "", fmt.Errorf("Unable to parse line %s", line)
			}
			return shortener.URL(splitString[1]), nil
		}
	}
	return "", fmt.Errorf("Pretend this is a nice custom error")
}
