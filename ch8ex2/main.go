package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/4ydx/cdp/protocol/dom"
	"github.com/4ydx/chrome-protocol"
	"github.com/4ydx/chrome-protocol/actions"
	"github.com/PuerkitoBio/goquery"
)

func getHTML() string {
	browser := cdp.NewBrowser("/usr/bin/google-chrome", 9222, "browser.log")
	handle := cdp.Start(browser, cdp.LogBasic)
	err := actions.EnableAll(handle, 2*time.Second)
	if err != nil {
		panic(err)
	}
	_, err = actions.Navigate(handle, "https://www.amazon.com/gp/goldbox", 30*time.Second)
	if err != nil {
		panic(err)
	}

	var nodes []dom.Node
	retries := 5

	for len(nodes) == 0 && retries > 0 {
		nodes, err = actions.FindAll(
			handle,
			"div.GB-M-COMMON.GB-SUPPLE:first-child #widgetContent",
			10*time.Second)
		retries--
		time.Sleep(1 * time.Second)
	}

	if len(nodes) == 0 || retries == 0 {
		panic("could not find results")
	}

	reply, err := actions.Evaluate(handle, "document.body.outerHTML;", 30*time.Second)
	if err != nil {
		panic(err)
	}

	a := struct{
		Value string
	}{}
	json.Unmarshal([]byte("{\"value\":" + string(*reply.Result.Value)+"}"), &a)
	body := a.Value

	handle.Stop(false)
	browser.Stop()
	return body
}

func parseProducts(htmlBody string) []string {
	println("parsing response")
	rdr := strings.NewReader(htmlBody)
	body, err := goquery.NewDocumentFromReader(rdr)
	if err != nil {
		panic(err)
	}

	products := []string{}
	details := body.Find("div.dealDetailContainer")
	println("Looking for products")
	details.Each(func(_ int, detail *goquery.Selection) {
		println(".")
		title := detail.Find("a#dealTitle").Text()
		price := detail.Find("div.priceBlock").Text()

		title = strings.TrimSpace(title)
		price = strings.TrimSpace(price)

		products = append(products, title + "\n"+price)
	})
	return products
}

func main() {
	println("getting HTML...")
	html := getHTML()
	println("parsing HTML...")
	products := parseProducts(html)

	println("Results:")
	for _, product := range products {
		fmt.Println(product + "\n")
	}
}
