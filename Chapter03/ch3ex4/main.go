package main

import (
	"io/ioutil"

	"github.com/gregjones/httpcache"
	"github.com/gregjones/httpcache/diskcache"
)

func main() {
	// Set up the local disk cache
	storage := diskcache.New("./cache")
	cache := httpcache.NewTransport(storage)

	// Set this to true to inform us if the responses are being read from a cache
	cache.MarkCachedResponses = true
	cachedClient := cache.Client()

	// Make the initial request
	println("Caching: http://www.example.com/index.html")
	resp, err := cachedClient.Get("http://www.example.com/index.html")
	if err != nil {
		panic(err)
	}

	// httpcache requires you to read the body in order to cache the response
	ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	// Request index.html again
	println("Requesting: http://www.example.com/index.html")
	resp, err = cachedClient.Get("http://www.example.com/index.html")
	if err != nil {
		panic(err)
	}

	// Look for the flag added by httpcache to show the result is read from the cache
	_, ok := resp.Header["X-From-Cache"]
	if ok {
		println("Result was pulled from the cache!")
	}
}
