package cli_test

import (
	"testing"
	"url-shortener/app"
	"url-shortener/cli"

	"github.com/stretchr/testify/assert"
)

func TestRun_store(t *testing.T) {
	s := cli.Run(app.NewInMemoryURLShortener(), "https://example.com/foo", "", false)
	assert.Equal(t, "Successfully shortened https://example.com/foo to Ti0-MV4cifgD\n", s)
}

func TestRun_storeAndRetrieve(t *testing.T) {
	s := cli.Run(app.NewInMemoryURLShortener(), "https://example.com/foo", "Ti0-MV4cifgD", false)
	assert.Equal(t, `Successfully shortened https://example.com/foo to Ti0-MV4cifgD
Successfully retrieved https://example.com/foo from Ti0-MV4cifgD
`, s)

	s = cli.Run(app.NewInMemoryURLShortener(), "https://example.com/foo", "unknown-code", false)
	assert.Equal(t, `Successfully shortened https://example.com/foo to Ti0-MV4cifgD
Unable to find unknown-code in the database
`, s)
}
