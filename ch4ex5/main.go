package main

import (
	"regexp"
	"strings"

	"github.com/antchfx/htmlquery"
)

func main() {
	doc, err := htmlquery.LoadURL("https://www.packtpub.com/packt/offers/free-learning")
	if err != nil {
		panic(err)
	}

	dealTextNodes := htmlquery.Find(doc, `//div[@class="dotd-main-book-summary float-left"]//text()`)

	if err != nil {
		panic(err)
	}

	println("Here is the free book of the day!")
	println("----------------------------------")

	for _, node := range dealTextNodes {
		text := strings.TrimSpace(node.Data)
		matchTagNames, _ := regexp.Compile("^(div|span|h2|br|ul|li)$")
		text = matchTagNames.ReplaceAllString(text,"")
		if text != "" {
			println(text)
		}
	}
}
