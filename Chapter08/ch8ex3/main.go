package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/slotix/dataflowkit/fetch"
)

func main() {
	r := fetch.Request{
		Type:      "base",
		URL:       "http://example.com",
		Method:    "GET",
		UserToken: "randomString",
		Actions:   "",
	}

	data, err := json.Marshal(&r)

	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:8000/fetch", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
