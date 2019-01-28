package main

import (
	"net/url"
	"path"
)

func main() {
	parsedUrl, err := url.Parse("https://hub.packtpub.com/8-programming-languages-to-learn-in-2019")

	if err != nil {
		panic(err)
	}

	site := parsedUrl.Host + parsedUrl.Path
	doesMatch, err := path.Match("hub.packtpub.com/*", site)
	if err != nil {
		panic(err)
	}
	if doesMatch {
		// Continue scraping …
		println("It's a match")
	}
}
