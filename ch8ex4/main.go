package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/slotix/dataflowkit/fetch"
	"github.com/slotix/dataflowkit/scrape"
)

func main() {
	r := scrape.Payload{
		Name: "Daily Deals",
		Request: fetch.Request{
			Type:   "Base",
			URL:    "https://www.packtpub.com/latest-releases",
			Method: "GET",
		},
		Fields: []scrape.Field{
			{
				Name:     "Title",
				Selector: `div.landing-page-row div[itemtype$="/Product"] div.book-block-title`,
				Extractor: scrape.Extractor{
					Types:   []string{"text"},
					Filters: []string{"trim"},
				},
			}, {
				Name:     "Price",
				Selector: `div.landing-page-row div[itemtype$="/Product"] div.book-block-price-discounted`,
				Extractor: scrape.Extractor{
					Types:   []string{"text"},
					Filters: []string{"trim"},
				},
			},
		},
		Format: "CSV",
	}

	data, err := json.Marshal(&r)

	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:8001/parse", "application/json", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
