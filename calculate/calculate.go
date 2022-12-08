package calculate

import (
	"fmt"

	"example.com/timeanddate"
	"example.com/wunder"
)

func CalcTemp(city string) int {
	fmt.Println("------------------ CalcTemp --------------------")
	// var city = "hadera"

	tempWunder := wunder.GetTemp(city)
	tempTimeAmdDate := timeanddate.GetTemp(city)
	fmt.Printf("By wunder temperature now at %s is: %d\n", city, tempWunder)
	fmt.Printf("By timeanddate temperature now at %s is: %d\n", city, tempTimeAmdDate)

	return (tempWunder + tempTimeAmdDate) / 2
}
