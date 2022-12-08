package calculate

import (
	"fmt"

	"example.com/atlas"
	"example.com/timeanddate"
	"example.com/wunder"
)

func CalcTemp(city string) int {
	fmt.Println("------------------ CalcTemp --------------------")
	// var city = "hadera"

	tempWunder := wunder.GetTemp(city)
	tempTimeAmdDate := timeanddate.GetTemp(city)
	tempAtlas := atlas.GetTemp(city)
	fmt.Printf("By wunderground temperature now at %s is: %d\n", city, tempWunder)
	fmt.Printf("By timeanddate temperature now at %s is: %d\n", city, tempTimeAmdDate)
	fmt.Printf("By weather-atlas temperature now at %s is: %d\n", city, tempAtlas)

	return (tempWunder + tempTimeAmdDate + tempAtlas) / 3
}
