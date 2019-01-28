package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Tracks the timestamp of the last request to the webserver
	var lastRequestMap map[string]time.Time = map[string]time.Time{
		"example.com":  time.Time{},
		"packtpub.com": time.Time{},
	}

	// The maximum number of requests we will make
	maximumNumberOfRequests := 5
	pageDelay := 5 * time.Second

	examplePage := "http://www.example.com/index.html"
	packtPage := "https://www.packtpub.com/"

	for i := 0; i < maximumNumberOfRequests; i++ {
		var elapsedTime time.Duration
		webpage := examplePage

		// Check if "i" is an even number
		if i%2 == 0 {
			// Use the Packt Publishing site and elapsed time
			webpage = packtPage
			elapsedTime = time.Now().Sub(lastRequestMap["packtpub.com"])
		} else {
			// Use the example.com elapsed time
			elapsedTime = time.Now().Sub(lastRequestMap["example.com"])
		}

		fmt.Printf("Elapsed Time: %.2f (s)\n", elapsedTime.Seconds())

		if elapsedTime < pageDelay {
			var timeDiff time.Duration = pageDelay - elapsedTime
			fmt.Printf("Sleeping for %.2f (s)\n", timeDiff.Seconds())
			time.Sleep(pageDelay - elapsedTime)
		}

		println("GET " + webpage)
		_, err := http.Get(webpage)
		if err != nil {
			panic(err)
		}

		// Update the last request time
		if i%2 == 0 {
			// Use the Packt Publishing elapsed time
			lastRequestMap["packtpub.com"] = time.Now()
		} else {
			// Use the example.com elapsed time
			lastRequestMap["example.com"] = time.Now()
		}
	}
}
