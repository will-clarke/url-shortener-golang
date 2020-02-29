package main

import (
	"flag"
	"fmt"
	"url-shortener/app"
	"url-shortener/cli"
	"url-shortener/hasher"
	"url-shortener/stores"
	"url-shortener/web"
)

func main() {
	useServer := flag.Bool("server", false, "start server")

	url := flag.String("shorten", "", "submit a URL To shorten & store")
	shortcode := flag.String("retrieve", "", "submit a code to retrieve URL from")
	redirect := flag.Bool("redirect", true, "should we redirect you?")

	flag.Parse()

	inMemoryShortener := newInMemoryURLShortener()
	if *useServer {
		web.Serve(inMemoryShortener, *redirect)
	}

	cli.Run(inMemoryShortener, *url, *shortcode, *redirect)

	if !*useServer && *url == "" && *shortcode == "" {
		helpMessage()
	}
}

func newInMemoryURLShortener() app.URLShortener {
	hashInstance := hasher.SHA256{}
	return app.URLShortener{
		Store:  stores.NewInMemoryStore(),
		Hasher: &hashInstance,
	}
}

func helpMessage() {
	fmt.Println("Run `-help` to see the options available")
}
