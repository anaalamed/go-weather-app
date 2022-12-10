package timeanddate

import (
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var url = "https://www.timeanddate.com/weather/israel/"

func GetTempToday(city string) int {
	c := initColly()
	log.Print("getTemp")

	var temp string
	c.OnHTML("#wt-ext tbody tr:first-child td:nth-child(5) ", func(e *colly.HTMLElement) {
		temp = strings.Fields(e.Text)[0]
	})

	c.Visit(url + city + "/ext")

	tempInt, err := strconv.Atoi(temp)
	if err != nil {
		log.Println("Error during conversion")
	}

	return tempInt
}

func GetAverageTemp(city string, days int) float32 {
	c := initColly()
	log.Print("getAverageTemp")

	var tempInts []int
	var tempStr string
	c.OnHTML("#wt-ext tbody tr td:nth-child(5)", func(e *colly.HTMLElement) {
		tempStr = strings.Fields(e.Text)[0]

		tempInt, err := strconv.Atoi(tempStr)
		if err != nil {
			log.Print("Error during conversion")
		}
		tempInts = append(tempInts, tempInt)
	})

	c.Visit(url + city + "/ext")

	tempInts = tempInts[:days]
	log.Printf("The temps for %d days are: %v", days, tempInts)

	return (getAverageArray(tempInts))
}

func GetTempMinMax(city string, days int) (int, int) {
	c := initColly()
	log.Print("GetTempMinMax")

	var tempInts []int
	var tempStr string
	c.OnHTML("#wt-ext tbody tr td:nth-child(5)", func(e *colly.HTMLElement) {
		tempStr = strings.Fields(e.Text)[0]

		tempInt, err := strconv.Atoi(tempStr)
		if err != nil {
			log.Print("Error during conversion")
		}
		tempInts = append(tempInts, tempInt)
	})

	c.Visit(url + city + "/ext")

	tempInts = tempInts[:days]
	min, max := getMinMaxArray(tempInts)

	log.Printf("The temps for %d days are: %v", days, tempInts)
	log.Printf("The min temp is: %d and the max temp is: %d ", min, max)

	return min, max
}

func initColly() *colly.Collector {
	log.SetPrefix("timeanddate: ")
	log.SetFlags(0)

	c := colly.NewCollector(
		colly.AllowedDomains("www.timeanddate.com", "timeanddate.com"),
	)

	c.OnResponse(func(r *colly.Response) {
		log.Print("Visited: ", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Print("Error!\n Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	return c
}

func getAverageArray(array []int) float32 {
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
