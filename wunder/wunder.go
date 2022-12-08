package wunder

import (
	"fmt"

	"github.com/gocolly/colly"
)

func Scrape() string {
	fmt.Println("Hello, World!")

	c := colly.NewCollector(
		colly.AllowedDomains("www.wunderground.com", "wunderground.com"),
	)

	// Called before a request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		fmt.Println("\n------------")
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error!\n Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Called after response received
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
		fmt.Println("\n------------")
	})

	var temperature string
	// Find  now temperature
	c.OnHTML(".station-nav .wu-value.wu-value-to", func(e *colly.HTMLElement) {
		temperature := e.Text
		fmt.Println(temperature)
	})

	c.Visit("https://www.wunderground.com/hourly/il/hadera")

	return temperature
}
