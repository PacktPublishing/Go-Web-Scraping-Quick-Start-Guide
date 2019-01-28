package main

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.packtpub.com/latest-releases")
	if err != nil {
		panic(err)
	}

	println("Here are the latest releases!")
	println("-----------------------------")
	doc.Find(`div.landing-page-row div[itemtype$="/Product"]`).
		Each(func(i int, e *goquery.Selection) {
			var title string
			var price float64
			
			title,_ = e.Attr("data-product-title")
			priceString, _ := e.Attr("data-product-price")
			price, err = strconv.ParseFloat(priceString, 64)
			if err != nil {
				println("Failed to parse price")
			}
			fmt.Printf("%s ($%0.2f)\n", title, price)
		})
}
