package main

import (
	"fmt"
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

	stringBody := string(data)

	numLinks := strings.Count(stringBody, "<a")
	fmt.Printf("Packt Publishing homepage has %d links!\n", numLinks)
}
