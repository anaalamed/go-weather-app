package atlas

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// var url = "https://www.weather-atlas.com/en/israel/"
var url = "https://www.weather-atlas.com/en/israel/"

func initColly() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("weather-atlas.com", "www.weather-atlas.com"),
	)

	return c
}

func GetTempToday(city string) int {
	log.SetPrefix("weather-atlas: ")
	log.SetFlags(0)
	log.Print("getTemp")

	c := initColly()

	c.OnResponse(func(r *colly.Response) {
		log.Print("Visited: ", r.Request.URL)
	})

	var temp string
	c.OnHTML(".card:nth-child(3) .card-body .row .row .fs-2", func(e *colly.HTMLElement) {
		temp = strings.Split(e.Text, "°")[0]
	})

	c.Visit(url + city + "-long-term-weather-forecast")

	tempInt, err := strconv.Atoi(temp)
	if err != nil {
		log.Print("Error during conversion")
	}

	return tempInt
}

func GetAverageTemp(city string, days int) float32 {
	log.SetPrefix("weather-atlas: ")
	log.SetFlags(0)
	log.Print("getAverageTemp")

	c := initColly()

	c.OnResponse(func(r *colly.Response) {
		log.Print("Visited: ", r.Request.URL)
	})

	var tempInts []int
	var tempStr string
	c.OnHTML(".row:nth-child(3) .card-body .fs-2.text-danger", func(e *colly.HTMLElement) {
		tempStr = strings.Split(e.Text, "°")[0]

		tempInt, err := strconv.Atoi(tempStr)
		if err != nil {
			log.Print("Error during conversion")
		}
		tempInts = append(tempInts, tempInt)
	})

	c.Visit(url + city + "-long-term-weather-forecast")
	tempInts = tempInts[:days]
	log.Printf("The temps for %d days are: %v\n", days, tempInts)

	return (averageArray(tempInts))
}

func averageArray(array []int) float32 {
	n := len(array)
	sum := 0

	for i := 0; i < n; i++ {
		sum += (array[i])
	}

	return (float32(sum) / float32(n))
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
