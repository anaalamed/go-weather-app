package timeanddate

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var url = "https://www.timeanddate.com/weather/israel/"

func initColly() *colly.Collector {

	c := colly.NewCollector(
		colly.AllowedDomains("www.timeanddate.com", "timeanddate.com"),
	)

	return c
}

func GetTemp(city string) int {
	log.SetPrefix("timeanddate: ")
	log.SetFlags(0)
	log.Print("getTemp")

	c := initColly()

	var temp string
	c.OnHTML("#wt-ext tbody tr:first-child td:nth-child(5) ", func(e *colly.HTMLElement) {
		temp = strings.Fields(e.Text)[0]
	})

	c.Visit(url + city + "/ext")

	tempInt, err := strconv.Atoi(temp)
	if err != nil {
		fmt.Println("Error during conversion")
	}

	return tempInt
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