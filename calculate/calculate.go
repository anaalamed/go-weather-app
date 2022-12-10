package calculate

import (
	"fmt"

	"example.com/atlas"
	"example.com/timeanddate"
	"example.com/wunder"
)

func CalcTemp(city string) float32 {
	fmt.Println("--------------------- Calculate Temperature Now ----------------------")
	// var city = "hadera"

	tempWunder := wunder.GetTemp(city)
	tempTimeAmdDate := timeanddate.GetTemp(city)
	tempAtlas := atlas.GetTemp(city)
	fmt.Printf("By wunderground temperature now at %s is: %d\n", city, tempWunder)
	fmt.Printf("By timeanddate temperature now at %s is: %d\n", city, tempTimeAmdDate)
	fmt.Printf("By weather-atlas temperature now at %s is: %d\n", city, tempAtlas)

	return average(tempAtlas, tempTimeAmdDate, tempWunder)
}

func average(num1 int, num2 int, num3 int) float32 {
	return (float32(num1) + float32(num2) + float32(num3)) / 3
}
