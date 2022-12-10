package wunder

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

// today only!
var url = "https://www.wunderground.com/forecast/il/"

func initColly() *colly.Collector {

	c := colly.NewCollector(
		colly.AllowedDomains("www.wunderground.com", "wunderground.com"),
	)

	return c
}

func GetTempToday(city string) int {
	log.SetPrefix("wunderground: ")
	log.SetFlags(0)
	log.Print("getTemp")

	c := initColly()

	c.OnResponse(func(r *colly.Response) {
		log.Print("Visited: ", r.Request.URL)
	})

	var temp string
	c.OnHTML(".station-nav .wu-value.wu-value-to", func(e *colly.HTMLElement) {
		// c.OnHTML(".city-almanac .content .row:nth-child(4) .columns:nth-child(3) span span", func(e *colly.HTMLElement) {
		temp = e.Text
	})

	c.Visit(url + city)

	// fahrenheit to celsius
	tempInt, err := strconv.Atoi(temp)

	if err != nil {
		fmt.Println("Error during conversion")
	}

	return (tempInt - 32) * 5 / 9
}

// func Scrape() string {
// 	fmt.Println("wunder")

// 	c := colly.NewCollector(
// 		colly.AllowedDomains("www.wunderground.com", "wunderground.com"),
// 	)

// 	// Called before a request
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 		fmt.Println("\n------------")
// 	})

// 	c.OnError(func(r *colly.Response, err error) {
// 		fmt.Println("Error!\n Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
// 	})

// 	// Called after response received
// 	c.OnResponse(func(r *colly.Response) {
// 		fmt.Println("Visited", r.Request.URL)
// 		fmt.Println("\n------------")
// 	})

// 	// Find now temperature
// 	var temperature = "1"
// 	c.OnHTML(".station-nav .wu-value.wu-value-to", func(e *colly.HTMLElement) {
// 		temperature := e.Text
// 		fmt.Println(temperature)
// 	})

// 	fmt.Println("after", temperature)

// 	c.Visit("https://www.wunderground.com/hourly/il/hadera")

// 	fmt.Println("end ", temperature)
// 	return temperature
// }
