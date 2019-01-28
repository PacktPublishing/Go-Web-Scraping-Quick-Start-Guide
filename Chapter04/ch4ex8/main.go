package main

import (
	"bufio"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	doc, err := goquery.NewDocument("https://www.packtpub.com/packt/offers/free-learning")
	if err != nil {
		panic(err)
	}

	println("Here is the free book of the day!")
	println("----------------------------------")
	rawText := doc.Find(`div.dotd-main-book-summary div:not(.eighteen-days-countdown-bar)`).Text()
	reader := bufio.NewReader(strings.NewReader(rawText))

	var line []byte
	for err == nil{
		line, _, err = reader.ReadLine()
		trimmedLine := strings.TrimSpace(string(line))
		if trimmedLine != "" {
			println(trimmedLine)
		}
	}
}
