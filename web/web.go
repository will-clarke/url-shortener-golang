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
	r := gin.Default()
	r.POST("/shorten", func(c *gin.Context) {
		type shortenRequest struct {
			URL string
		}
		var r shortenRequest
		c.Bind(&r)
		if r.URL == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": "false",
				"error":   "no 'url' in POST body found",
			})
			return
		}
		shortCode, err := urlShortener.StoreURL(shortener.URL(r.URL))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": "false",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success":     true,
			"redirectURL": HOST + string(shortCode),
			"error":       nil,
		})
	})
	r.GET("/:code", func(c *gin.Context) {
		code := c.Param("code")
		shortCode := shortener.ShortCode(code)
		url, err := urlShortener.GetURL(shortCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": "false",
				"error":   err.Error(),
			})
			return
		}
		if url == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"success": "false",
				"error":   fmt.Sprint("code '", code, "' not found in database"),
			})
			return
		}
		if redirect {
			c.Redirect(http.StatusMovedPermanently, string(url))
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success":     true,
			"error":       nil,
			"redirectURL": string(url),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
