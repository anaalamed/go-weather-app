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

	return average(tempAtlas, tempTimeAmdDate, tempWunder)
}

func CalcAverageTemp(city string, days int) float32 {
	fmt.Println("\n\n--------------------- Calculate Average Temperature for days ----------------------")

	tempTimeAmdDate := timeanddate.GetAverageTemp(city, days)
	tempAtlas := atlas.GetAverageTemp(city, days)
	fmt.Printf("\nBy timeanddate temperature for %d days at %s is: %.2f\n", days, city, tempTimeAmdDate)
	fmt.Printf("By weather-atlas temperature for %d days at %s is: %.2f\n", days, city, tempAtlas)

	return averageFloat(tempAtlas, tempTimeAmdDate)
}

func average(num1 int, num2 int, num3 int) float32 {
	return (float32(num1) + float32(num2) + float32(num3)) / 3
}

func averageFloat(num1 float32, num2 float32) float32 {
	return (num1 + num2) / 2
}
