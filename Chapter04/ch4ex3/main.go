package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
        "regexp"
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

        re := regexp.MustCompile(`<a.*href\s*=\s*["'](http[s]{0,1}:\/\/.[^\s]*)["'].*>`)
        linkMatches := re.FindAllStringSubmatch(stringBody, -1)

        fmt.Printf("Found %d links:\n", len(linkMatches))
        for _,linkGroup := range(linkMatches){
            println(linkGroup[1])
        }
}
