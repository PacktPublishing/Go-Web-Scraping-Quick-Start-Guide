package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.packtpub.com/latest-releases")
	if err != nil {
		panic(err)
	}

	println("Here are the latest releases!")
	println("-----------------------------")
	time.Sleep(1 * time.Second)
	doc.Find(`div.landing-page-row div[itemtype$="/Product"] a`).
		Each(func(i int, e *goquery.Selection) {
			var title, description, author, price string
			link, _ := e.Attr("href")
			link = "https://www.packtpub.com" + link

			bookPage, err := goquery.NewDocument(link)
			if err != nil {
				panic(err)
			}
			title = bookPage.Find("div.book-top-block-info h1").Text()
			description = strings.TrimSpace(bookPage.Find("div.book-top-block-info div.book-top-block-info-one-liner").Text())
			price = strings.TrimSpace(bookPage.Find("div.book-top-block-info div.onlyDesktop div.book-top-pricing-main-ebook-price").Text())
			authorNodes := bookPage.Find("div.book-top-block-info div.book-top-block-info-authors")
			if len(authorNodes.Nodes) < 1 {
				return
			}
			author = strings.TrimSpace(authorNodes.Nodes[0].FirstChild.Data)
			fmt.Printf("%s\nby: %s\n%s\n%s\n---------------------\n\n", title, author, price, description)
			time.Sleep(1 * time.Second)
		})
}
