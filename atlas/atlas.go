package atlas

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var url = "https://www.weather-atlas.com/en/israel/"

func initColly() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("weather-atlas.com", "www.weather-atlas.com"),
	)

	return c
}

func GetTemp(city string) int {
	log.SetPrefix("weather-atlas: ")
	log.SetFlags(0)
	log.Print("getTemp")

	c := initColly()

	var temp string
	c.OnHTML(".card-body .row .row .fs-2", func(e *colly.HTMLElement) {
		temp = strings.Split(e.Text, "°")[0]
	})

	c.Visit(url + city)

	tempInt, err := strconv.Atoi(temp)
	if err != nil {
		log.Print("Error during conversion")
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
