package cli

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"url-shortener/app"
	"url-shortener/shortener"
)

func Run(urlShortener app.URLShortener, urlString, shortCodeString string, redirect bool) (s string) {
	if urlString != "" {
		s += storeURL(urlShortener, urlString)
	}

	if shortCodeString != "" {
		s += getURL(urlShortener, shortCodeString, redirect)
	}
	return s
}

func storeURL(urlShortener app.URLShortener, urlString string) string {
	url := shortener.URL(urlString)
	code, err := urlShortener.StoreURL(url)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintln("Successfully shortened", urlString, "to", string(code))
}

func getURL(urlShortener app.URLShortener, shortCodeString string, redirect bool) string {
	shortCode := shortener.ShortCode(shortCodeString)
	url, err := urlShortener.GetURL(shortCode)
	if err != nil {
		return fmt.Sprintln(err)
	}
	if string(url) == "" {
		return fmt.Sprintln("Unable to find", shortCodeString, "in the database")
	}
	if redirect {
		openBrowser(string(url))
	}
	return fmt.Sprintln("Successfully retrieved", string(url), "from", shortCodeString)
}

// cribbed from https://gist.github.com/nanmu42/4fbaf26c771da58095fa7a9f14f23d27
func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
