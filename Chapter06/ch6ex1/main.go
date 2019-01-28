package main

import (
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// Public proxies from https://hidemyna.me
var proxies []string = []string{
	"http://207.154.231.208:8080",
	"http://138.68.230.88:8080",
	"http://162.243.107.45:8080",
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
		resp, err := http.Get("http://ip-api.com/line")
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
