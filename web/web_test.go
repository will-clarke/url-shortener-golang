package web_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"git.sr.ht/~will-clarke/url-shortener-golang/app"
	"git.sr.ht/~will-clarke/url-shortener-golang/web"

	"github.com/stretchr/testify/assert"
)

func TestServeGetNonExisting(t *testing.T) {
	inMemoryShortener := app.NewInMemoryURLShortener()
	router := web.SetupRouter(inMemoryShortener, false)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/not-exist-in-db", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	expectedBody := `{"success":false,"error":"code 'not-exist-in-db' not found in database","redirectURL":""}`
	assert.Equal(t, expectedBody, w.Body.String())
}

func TestStore(t *testing.T) {
	inMemoryShortener := app.NewInMemoryURLShortener()
	router := web.SetupRouter(inMemoryShortener, false)

	w := httptest.NewRecorder()
	postJSON := `{"URL": "https://example.com/foo"}`
	req, _ := http.NewRequest("POST", "/shorten", strings.NewReader(postJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	expectedBody := `{"success":true,"redirectURL":"localhost:8080/Ti0-MV4cifgD"}`
	assert.Equal(t, expectedBody, w.Body.String())
}

func TestStoreAndRedirect(t *testing.T) {
	inMemoryShortener := app.NewInMemoryURLShortener()
	router := web.SetupRouter(inMemoryShortener, true)

	// test the POST /shorten
	w := httptest.NewRecorder()
	postJSON := `{"URL": "https://example.com/foo"}`
	req, _ := http.NewRequest("POST", "/shorten", strings.NewReader(postJSON))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	expectedBody := `{"success":true,"redirectURL":"localhost:8080/Ti0-MV4cifgD"}`
	assert.Equal(t, expectedBody, w.Body.String())

	// test the GET /:code
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/Ti0-MV4cifgD", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 301, w.Code)
	expectedBody = "<a href=\"https://example.com/foo\">Moved Permanently</a>.\n\n"
	assert.Equal(t, expectedBody, w.Body.String())
}
