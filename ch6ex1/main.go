package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// Public proxies from https://free-proxy-list.net/
var proxies []string = []string{
	"http://76.123.162.68:8080",
	"http://162.223.89.92:8080",
	"http://104.248.183.73:80",
}

func GetProxy(_ *http.Request) (*url.URL, error) {
	randomIndex := rand.Int31n(int32(len(proxies)))
	randomProxy := proxies[randomIndex]

	return url.Parse(randomProxy)
}

func main() {
	rand.Seed(time.Now().Unix())
	http.DefaultTransport.(*http.Transport).Proxy = GetProxy

	// Continue with your HTTP requests
	for i := 0; i < 5; i++ {
		resp, err := http.Get("https://api.ipify.org")
		if err != nil {
			panic(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		println("Proxy IP is: " + string(data))
		time.Sleep(1 * time.Second)
	}
}
