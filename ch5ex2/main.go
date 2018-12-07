package main

import (
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{}
	data.Set("s", "Golang")

	response, err := http.PostForm("https://hub.packtpub.com/", data)
	
	// ... Continue processing the response ...
}
