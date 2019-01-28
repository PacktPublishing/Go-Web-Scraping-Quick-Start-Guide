package main

import (
	"github.com/tebeka/selenium"
)

func main() {

	// The paths to these binaries will be different on your machine!

	const (
		seleniumPath    = "/home/vincent/Documents/workspace/Go/src/github.com/tebeka/selenium/vendor/selenium-server-standalone-3.14.0.jar"
		geckoDriverPath = "/home/vincent/Documents/workspace/Go/src/github.com/tebeka/selenium/vendor/geckodriver-v0.23.0-linux64"
	)

	service, err := selenium.NewSeleniumService(
		seleniumPath,
		8080,
		selenium.GeckoDriver(geckoDriverPath))

	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, "http://localhost:8080/wd/hub")
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	err = wd.Get("https://www.packtpub.com/networking-and-servers/mastering-go")
	if err != nil {
		panic(err)
	}

	var elems []selenium.WebElement
	wd.Wait(func(wd2 selenium.WebDriver) (bool, error) {
		elems, err = wd.FindElements(selenium.ByCSSSelector, "div.product-reviews-review div.review-body")
		if err != nil {
			return false, err
		} else {
			return len(elems) > 0, nil
		}
	})

	for _, review := range elems {
		body, err := review.Text()
		if err != nil {
			panic(err)
		}
		println(body)
	}
}
