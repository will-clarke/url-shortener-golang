package cli

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"url-shortener/app"
	"url-shortener/shortener"
)

func Run(urlShortener app.URLShortener, urlString, shortCodeString string, redirect bool) {
	if urlString != "" {
		storeURL(urlShortener, urlString)
	}

	if shortCodeString != "" {
		getURL(urlShortener, shortCodeString, redirect)
	}
}

func storeURL(urlShortener app.URLShortener, urlString string) {
	url := shortener.URL(urlString)
	code, err := urlShortener.StoreURL(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully shortened", urlString, "to", string(code))
}

func getURL(urlShortener app.URLShortener, shortCodeString string, redirect bool) {
	shortCode := shortener.ShortCode(shortCodeString)
	url, err := urlShortener.GetURL(shortCode)
	if err != nil {
		fmt.Println(err)
	}
	if string(url) == "" {
		fmt.Println("Unable to find", shortCodeString, "in the database")
		return
	}
	fmt.Println("Successfully retrieved", string(url), "from", shortCodeString)
	if redirect {
		openBrowser(string(url))
	}
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
