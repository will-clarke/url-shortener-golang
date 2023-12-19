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

	// Because we used an interface, swapping out the backend stores is pretty easy:
	// store := app.NewInMemoryURLShortener()
	store, err := app.NewFileStore()
	if err != nil {
		fmt.Println("Error creating filestore", err)
		panic(err)
	}

	if *useServer {
		web.Serve(store, *redirect)
	}

	cliOutput := cli.Run(store, *url, *shortcode, *redirect)
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
