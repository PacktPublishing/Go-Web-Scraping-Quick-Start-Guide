package main

import (
	"fmt"
	"strconv"

	"github.com/antchfx/htmlquery"
)

func main() {
	doc, err := htmlquery.LoadURL("https://www.packtpub.com/latest-releases")
	if err != nil {
		panic(err)
	}

	nodes := htmlquery.Find(doc, `//div[@class="landing-page-row cf"]/div[@itemtype="http://schema.org/Product"]`)
	if err != nil {
		panic(err)
	}

	println("Here are the latest releases!")
	println("-----------------------------")

	for _, node := range nodes {
		var title string
		var price float64

		for _, attribute := range node.Attr {
			switch attribute.Key {
			case "data-product-title":
				title = attribute.Val
			case "data-product-price":
				price, err = strconv.ParseFloat(attribute.Val, 64)
				if err != nil {
					println("Failed to parse price")
				}
			}
		}
		fmt.Printf("%s ($%0.2f)\n", title, price)
	}
}
