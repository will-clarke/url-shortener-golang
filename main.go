package main

import (
	"flag"
	"fmt"
	"url-shortener/app"
	"url-shortener/cli"
	"url-shortener/hasher"
	"url-shortener/stores"
)

func main() {
	useServer := flag.Bool("server", false, "start server")

	url := flag.String("shorten", "", "submit a URL To shorten & store")
	shortcode := flag.String("retrieve", "", "submit a code to retrieve URL from")
	redirect := flag.Bool("redirect", false, "should the CLI open a browser & redirect you?")

	flag.Parse()

	inMemoryShortener := newInMemoryURLShortener()
	if *useServer {
		fmt.Println("using server")
		fmt.Println("using server")
		fmt.Println("using server")
		fmt.Println("using server")
		fmt.Println("using server")
		return
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
