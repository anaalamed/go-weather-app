package calculate

import (
	"fmt"

	"example.com/atlas"
	"example.com/timeanddate"
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
	return getAverageArrayInt(tempArr)
}

func CalcAverageTemp(city string, days int) float32 {
	fmt.Println("\n\n--------------------- Calculate Average Temperature for days ----------------------")

	tempTimeAmdDate := timeanddate.GetAverageTemp(city, days)
	tempAtlas := atlas.GetAverageTemp(city, days)
	fmt.Printf("\nBy timeanddate temperature for %d days at %s is: %.2f\n", days, city, tempTimeAmdDate)
	fmt.Printf("By weather-atlas temperature for %d days at %s is: %.2f\n", days, city, tempAtlas)

	tempArr := []float32{tempAtlas, tempTimeAmdDate}
	return getAverageArrayFloat(tempArr)
}

func GetTempMinMax(city string, days int) (float32, float32) {
	fmt.Println("\n\n--------------------- Calculate Average Min and Max Temperature for days ----------------------")

	tempTimeAmdDateMin, tempTimeAmdDateMax := timeanddate.GetTempMinMax(city, days)
	tempAtlasMin, tempAtlasMax := atlas.GetTempMinMax(city, days)
	fmt.Printf("\nBy timeanddate for %d days at %s min temp is: %d and max temp is: %d\n", days, city, tempTimeAmdDateMin, tempTimeAmdDateMax)
	fmt.Printf("By weather-atlas for %d days at %s min temp is: %d and max temp is: %d\n", days, city, tempAtlasMin, tempAtlasMax)

	minArr := []int{tempAtlasMin, tempTimeAmdDateMin}
	maxArr := []int{tempAtlasMax, tempTimeAmdDateMax}
	return getAverageArrayInt(minArr), getAverageArrayInt(maxArr)
}

func getAverageArrayInt(array []int) float32 {
	n := len(array)
	sum := 0

	for i := 0; i < n; i++ {
		sum += (array[i])
	}

	return (float32(sum) / float32(n))
}

func getAverageArrayFloat(array []float32) float32 {
	n := len(array)
	var sum float32 = 0

	for i := 0; i < n; i++ {
		sum += (array[i])
	}

	return (sum / float32(n))
}
