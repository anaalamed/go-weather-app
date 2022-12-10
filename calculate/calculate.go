package calculate

import (
	"fmt"

	"example.com/atlas"
	"example.com/timeanddate"
	"example.com/utils"
	"example.com/wunder"
)

func CalcTempToday(city string) float32 {
	fmt.Println("--------------------- Calculate Temperature Now ----------------------")

	tempWunder := wunder.GetTempToday(city)
	tempTimeAmdDate := timeanddate.GetTempToday(city)
	tempAtlas := atlas.GetTempToday(city)
	fmt.Printf("\nBy wunderground temperature now at %s is: %d\n", city, tempWunder)
	fmt.Printf("By timeanddate temperature now at %s is: %d\n", city, tempTimeAmdDate)
	fmt.Printf("By weather-atlas temperature now at %s is: %d\n", city, tempAtlas)

	tempArr := []int{tempAtlas, tempTimeAmdDate, tempWunder}
	return utils.GetAverageArrayInt(tempArr)
}

func CalcAverageTemp(city string, days int) float32 {
	fmt.Println("\n\n--------------------- Calculate Average Temperature for days ----------------------")

	tempTimeAmdDate := timeanddate.GetAverageTemp(city, days)
	tempAtlas := atlas.GetAverageTemp(city, days)
	fmt.Printf("\nBy timeanddate temperature for %d days at %s is: %.2f\n", days, city, tempTimeAmdDate)
	fmt.Printf("By weather-atlas temperature for %d days at %s is: %.2f\n", days, city, tempAtlas)

	tempArr := []float32{tempAtlas, tempTimeAmdDate}
	return utils.GetAverageArrayFloat(tempArr)
}

func CalcTempMinMax(city string, days int) (float32, float32) {
	fmt.Println("\n\n--------------------- Calculate Average Min and Max Temperature for days ----------------------")

	tempTimeAmdDateMin, tempTimeAmdDateMax := timeanddate.GetTempMinMax(city, days)
	tempAtlasMin, tempAtlasMax := atlas.GetTempMinMax(city, days)
	fmt.Printf("\nBy timeanddate for %d days at %s min temp is: %d and max temp is: %d\n", days, city, tempTimeAmdDateMin, tempTimeAmdDateMax)
	fmt.Printf("By weather-atlas for %d days at %s min temp is: %d and max temp is: %d\n", days, city, tempAtlasMin, tempAtlasMax)

	minArr := []int{tempAtlasMin, tempTimeAmdDateMin}
	maxArr := []int{tempAtlasMax, tempTimeAmdDateMax}
	return utils.GetAverageArrayInt(minArr), utils.GetAverageArrayInt(maxArr)
}

func CalcWeatherSummary(city string) *utils.Weather {
	fmt.Println("\n\n--------------------- Calculate Weather Summary ----------------------")

	weatherWunder := wunder.GetWeatherSummary(city)
	weatherAtlas := atlas.GetWeatherSummary(city)
	weatherTimeAndDate := timeanddate.GetWeatherSummary(city)

	fmt.Printf("\nBy wunderground the weather summary is: %+v\n", weatherWunder)
	fmt.Printf("By atlas-weather the weather summary is: %+v\n", weatherAtlas)
	fmt.Printf("By timeanddate the weather summary is: %+v\n", weatherTimeAndDate)

	tempAvg := utils.GetAverageArrayFloat([]float32{weatherWunder.Temp, weatherAtlas.Temp, weatherTimeAndDate.Temp})
	humidityAvg := utils.GetAverageArrayFloat([]float32{weatherWunder.Humidity, weatherAtlas.Humidity, weatherTimeAndDate.Humidity})
	precipitationAvg := utils.GetAverageArrayFloat([]float32{weatherWunder.Precipitation, weatherAtlas.Precipitation, weatherTimeAndDate.Precipitation})

	weatherSummaryAvg := &utils.Weather{
		Temp:          tempAvg,
		Humidity:      humidityAvg,
		Precipitation: precipitationAvg,
	}
	return weatherSummaryAvg
}
