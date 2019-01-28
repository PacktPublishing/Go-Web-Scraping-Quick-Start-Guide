package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	resp, err := http.Get("https://www.packtpub.com/")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	stringBody := strings.ToLower(string(data))

	if strings.Contains(stringBody, "<!doctype html>") {
		println("This webpage is HTML5")
	} else if strings.Contains(stringBody, "html/strict.dtd") {
		println("This webpage is HTML4 (Strict)")
	} else if strings.Contains(stringBody, "html/loose.dtd") {
		println("This webpage is HTML4 (Tranistional)")
	} else if strings.Contains(stringBody, "html/frameset.dtd") {
		println("This webpage is HTML4 (Frameset)")
	} else {
		println("Could not determine doctype!")
	}
}
