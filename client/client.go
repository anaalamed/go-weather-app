package main

import (
	"fmt"

	"example.com/calculate"
)

func main() {
	fmt.Println("Hello, boooooooom!")
	var city = "hadera"

	tempToday := calculate.CalcTempToday(city)
	fmt.Printf("\nThe average temperature Today at %s is: %.2f\n", city, tempToday)

	var days = 5
	temp := calculate.CalcAverageTemp(city, days)
	fmt.Printf("\nThe average temperature for %d days at %s is: %.2f\n", days, city, temp)
}
