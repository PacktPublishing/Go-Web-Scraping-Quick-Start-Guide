package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Tracks the timestamp of the last request to the webserver
	var lastRequestTime time.Time

	// The maximum number of requests we will make to the webserver
	maximumNumberOfRequests := 5

	// Our scrape rate at 1 page per 5 seconds
	pageDelay := 5 * time.Second

	for i := 0; i < maximumNumberOfRequests; i++ {
		// Calculate the time difference since our last request
		elapsedTime := time.Now().Sub(lastRequestTime)
		fmt.Printf("Elapsed Time: %.2f (s)\n", elapsedTime.Seconds())

		//Check if there has been enough time
		if elapsedTime < pageDelay {
			// Sleep the difference between the pageDelay and elapsedTime
			var timeDiff time.Duration = pageDelay - elapsedTime
			fmt.Printf("Sleeping for %.2f (s)\n", timeDiff.Seconds())
			time.Sleep(pageDelay - elapsedTime)
		}

		// Just for this example, we are not processing the response
		println("GET example.com/index.html")
		_, err := http.Get("http://www.example.com/index.html")
		if err != nil {
			panic(err)
		}

		// Update the last request time
		lastRequestTime = time.Now()
	}
}
