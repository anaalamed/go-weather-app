package atlas

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var url = "https://www.weather-atlas.com/en/israel/"

func GetTempToday(city string) int {
	c := initColly()
	log.Print("getTemp")

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
	c := initColly()
	log.Print("getAverageTemp")

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

func GetTempMinMax(city string, days int) (int, int) {
	c := initColly()
	log.Print("GetTempMinMax")

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
	min, max := getMinMaxArray(tempInts)

	log.Printf("The temps for %d days are: %v", days, tempInts)
	log.Printf("The min temp is: %d and the max temp is: %d ", min, max)

	return min, max
}

func initColly() *colly.Collector {
	log.SetPrefix("weather-atlas: ")
	log.SetFlags(0)

	c := colly.NewCollector(
		colly.AllowedDomains("weather-atlas.com", "www.weather-atlas.com"),
	)

	c.OnResponse(func(r *colly.Response) {
		log.Print("Visited: ", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Print("Error!\n Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	return c
}

func averageArray(array []int) float32 {
	n := len(array)
	sum := 0

	for i := 0; i < n; i++ {
		sum += (array[i])
	}

	return (float32(sum) / float32(n))
}

func getMinMaxArray(array []int) (int, int) {
	min := 1000
	max := 0
	for i := 0; i < len(array); i++ {
		if max < array[i] {
			max = array[i]
		} else if min > array[i] {
			min = array[i]
		}
	}

	return min, max
}
