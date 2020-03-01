package web

import (
	"fmt"
	"net/http"
	"url-shortener/app"
	"url-shortener/shortener"

	"github.com/gin-gonic/gin"
)

const HOST = "localhost:8080/" // probs use an ENV VAR for this

func Serve(urlShortener app.URLShortener, redirect bool) {
	r := SetupRouter(urlShortener, redirect)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func SetupRouter(urlShortener app.URLShortener, redirect bool) *gin.Engine {
	r := gin.Default()
	h := handler{
		urlShortener: urlShortener,
		redirect:     redirect,
	}
	r.POST("/shorten", h.shorten)
	r.GET("/:code", h.getURL)
	return r
}

type response struct {
	Success     bool   `json:"success"`
	Error       string `json:"error",omitempty:true`
	RedirectURL string `json:"redirectURL"`
}

type handler struct {
	urlShortener app.URLShortener
	redirect     bool
}

func (h *handler) shorten(c *gin.Context) {
	type shortenRequest struct {
		URL string
	}
	var r shortenRequest
	err := c.Bind(&r)
	if r.URL == "" {
		c.JSON(http.StatusBadRequest,
			response{
				Success: false,
				Error:   "no 'url' in POST body found",
			})
		return
	}
	shortCode, err := h.urlShortener.StoreURL(shortener.URL(r.URL))
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			response{
				Success: false,
				Error:   err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK, response{
		Success:     true,
		RedirectURL: HOST + string(shortCode),
	})
}

func (h *handler) getURL(c *gin.Context) {
	code := c.Param("code")
	shortCode := shortener.ShortCode(code)
	url, err := h.urlShortener.GetURL(shortCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response{
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	if url == "" {
		c.JSON(http.StatusNotFound, response{
			Success: false,
			Error:   fmt.Sprint("code '", code, "' not found in database"),
		})
		return
	}
	if h.redirect {
		c.Redirect(http.StatusMovedPermanently, string(url))
		return
	}
	c.JSON(http.StatusOK, response{
		Success:     true,
		RedirectURL: string(url),
	})
}
