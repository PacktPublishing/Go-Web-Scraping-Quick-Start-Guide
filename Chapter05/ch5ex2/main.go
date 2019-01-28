package main

import (
	"net/http"
	"net/url"
)

func main() {
	data := url.Values{}
	data.Set("s", "Golang")

	response, err := http.PostForm("https://hub.packtpub.com/", data)
	if err != nil {
		panic(err)
	}
	// ... Continue processing the response ...
	println(response.StatusCode)
}
