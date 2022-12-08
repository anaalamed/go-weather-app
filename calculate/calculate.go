package calculate

import (
	"fmt"

	"example.com/wunder"
)

func CalcTemp(city string) int {
	fmt.Println("------------------ CalcTemp --------------------")
	// var city = "hadera"

	tempWunder := wunder.GetTemp(city)
	fmt.Printf("By wunder temperature now at %s is: %d\n", city, tempWunder)
	return tempWunder
}
