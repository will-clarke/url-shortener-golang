package main

import (
	"flag"
	"fmt"

	"git.sr.ht/~will-clarke/url-shortner-golang/app"
	"git.sr.ht/~will-clarke/url-shortner-golang/cli"
	"git.sr.ht/~will-clarke/url-shortner-golang/web"
)

func main() {
	useServer := flag.Bool("server", false, "start server")

	url := flag.String("shorten", "", "submit a URL To shorten & store")
	shortcode := flag.String("retrieve", "", "submit a code to retrieve URL from")
	redirect := flag.Bool("redirect", true, "should we redirect you?")

	flag.Parse()

	inMemoryShortener := app.NewInMemoryURLShortener()
	if *useServer {
		web.Serve(inMemoryShortener, *redirect)
	}

	cliOutput := cli.Run(inMemoryShortener, *url, *shortcode, *redirect)
	if cliOutput != "" {
		fmt.Print(cliOutput)
	}

	if !*useServer && *url == "" && *shortcode == "" {
		helpMessage()
	}
}

func helpMessage() {
	fmt.Println("Run `-help` to see the options available")
}
