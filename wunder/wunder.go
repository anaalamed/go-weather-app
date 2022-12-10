package wunder

import (
	"log"
	"strconv"

	"example.com/utils"

	"github.com/gocolly/colly"
)

// today only!
var url = "https://www.wunderground.com/forecast/il/"

func initColly() *colly.Collector {
	log.SetPrefix("wunderground: ")
	log.SetFlags(0)

	c := colly.NewCollector(
		colly.AllowedDomains("www.wunderground.com", "wunderground.com"),
	)

	c.OnResponse(func(r *colly.Response) {
		log.Print("Visited: ", r.Request.URL)
	})

	return c
}

func GetTempToday(city string) int {
	c := initColly()
	log.Print("getTemp")

	var temp string
	c.OnHTML(".station-nav .wu-value.wu-value-to", func(e *colly.HTMLElement) {
		// c.OnHTML(".city-almanac .content .row:nth-child(4) .columns:nth-child(3) span span", func(e *colly.HTMLElement) {
		temp = e.Text
	})

	c.Visit(url + city)

	// fahrenheit to celsius
	tempInt, err := strconv.Atoi(temp)

	if err != nil {
		log.Println("Error during conversion")
	}

	return (tempInt - 32) * 5 / 9
}

func GetWeatherSummary(city string) *utils.Weather {
	c := initColly()
	log.Print("GetWeatherSummary")

	var temp, humidity, precipitation string

	c.OnHTML(".station-nav .wu-value.wu-value-to", func(e *colly.HTMLElement) {
		temp = e.Text
	})

	c.OnHTML(".wu-unit.wu-unit-humidity .wu-value.wu-value-to", func(e *colly.HTMLElement) {
		humidity = e.Text
	})

	c.OnHTML(".wu-unit.wu-unit-rain .wu-value.wu-value-to", func(e *colly.HTMLElement) {
		precipitation = e.Text
	})

	c.Visit(url + city)

	tempInt, err := strconv.Atoi(temp)
	humidityInt, err := strconv.Atoi(humidity)
	precipitationInt, err := strconv.Atoi(precipitation)

	if err != nil {
		log.Println("Error during conversion")
	}

	weatherSummary := &utils.Weather{
		Temp:          float32((tempInt - 32) * 5 / 9),
		Humidity:      float32(humidityInt),
		Precipitation: float32(precipitationInt),
	}
	return weatherSummary
}
