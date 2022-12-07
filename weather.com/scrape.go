package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// forbidden
	fmt.Println("accuweather")

	c := colly.NewCollector(
		colly.AllowedDomains("accuweather.com"),
	)

	// Called before a request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		fmt.Println("\n------------")
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error - Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Called after response received
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		fmt.Println("\n------------")
	})

	// Called right after OnResponse if the received content is HTML
	c.OnHTML(".detail-item", func(e *colly.HTMLElement) {
		// links := e.ChildAttrs("a", "href")
		// fmt.Println(links)
		// divs := e.Text
		fmt.Println("boom", e.Text)
	})
	c.Visit("https://accuweather.com/en/il/hadera/213124/current-weather/213124")
}
