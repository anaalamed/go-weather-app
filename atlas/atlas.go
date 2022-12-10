package atlas

import (
	"log"
	"strconv"
	"strings"

	"example.com/utils"
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

	return (utils.GetAverageArrayInt(tempInts))
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
	min, max := utils.GetMinMaxArray(tempInts)

	log.Printf("The temps for %d days are: %v", days, tempInts)
	log.Printf("The min temp is: %d and the max temp is: %d ", min, max)

	return min, max
}

func GetWeatherSummary(city string) *utils.Weather {
	c := initColly()
	log.Print("GetWeatherSummary")

	var temp, humidity, precipitation string

	c.OnHTML(".card:nth-child(3) .card-body .row .row .fs-2", func(e *colly.HTMLElement) {
		temp = strings.Split(e.Text, "°")[0]
	})

	c.OnHTML(".card:nth-child(3) .card-body .row .row div:nth-child(2) ul li:nth-child(2)", func(e *colly.HTMLElement) {
		humidity = strings.Split(strings.Fields(e.Text)[1], "%")[0]
	})

	c.OnHTML(".card:nth-child(3) .card-body .row .row div:nth-child(2) ul li:nth-child(4)", func(e *colly.HTMLElement) {
		precipitation = strings.Split(strings.Fields(e.Text)[1], "mm")[0]
	})

	c.Visit(url + city + "-long-term-weather-forecast")

	tempInt, err := strconv.Atoi(temp)
	humidityInt, err := strconv.Atoi(humidity)
	precipitationInt, err := strconv.Atoi(precipitation)

	if err != nil {
		log.Println("Error during conversion")
	}

	weatherSummary := &utils.Weather{
		Temp:          float32(tempInt),
		Humidity:      float32(humidityInt),
		Precipitation: float32(precipitationInt),
	}
	return weatherSummary
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
