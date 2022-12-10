package main

import (
	"fmt"

	"example.com/calculate"
)

func main() {
	fmt.Println("Hello, boooooooom!")
	var city = "hadera" // entire name only
	var days = 9        // <9

	tempToday := calculate.CalcTempToday(city)
	fmt.Printf("\nThe average temperature Today at %s is: %.2f\n", city, tempToday)

	temp := calculate.CalcAverageTemp(city, days)
	fmt.Printf("\nThe average temperature for %d days at %s is: %.2f\n", days, city, temp)

	tempMin, tempMax := calculate.GetTempMinMax(city, days)
	fmt.Printf("\nFor %d days at %s the min temperature is: %.2f and the max temperature is: %.2f\n", days, city, tempMin, tempMax)

}
