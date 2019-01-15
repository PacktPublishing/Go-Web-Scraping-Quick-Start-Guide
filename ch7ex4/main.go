package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var sites []string = []string{
	"http://example.com/site1.html",
	"http://example.com/site2.html",
	"http://example.com/site3.html",
}
var activeThreads int32 = 0
var doneCount = 0

const maxActiveThreads = 1

func scrapeSite(site string, condition *sync.Cond) {
	condition.L.Lock()
	if activeThreads >= maxActiveThreads {
		println(site + " Max threads reached")
		condition.Wait()
	}
	condition.L.Unlock()

	// activeThreads = atomic.AddInt32(&activeThreads, 1)
	atomic.AddInt32(&activeThreads, 1)
	// Scraping code goes here ...
	println("scraping " + site)
	// activeThreads = atomic.AddInt32(&activeThreads, -1)
	atomic.AddInt32(&activeThreads, -1)
	condition.Signal()
}

func main() {
	var l = sync.Mutex{}
	var c = sync.NewCond(&l)

	for _, site := range sites {
		println("starting scraper for " + site)
		go scrapeSite(site, c)
	}
	for doneCount < len(sites) {
		time.Sleep(1 * time.Second)
	}
	println("Done!")
}
