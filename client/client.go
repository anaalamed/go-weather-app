package main

import (
	"fmt"

	"example.com/calculate"
)

func main() {
	fmt.Println("Hello, boooooooom!")
	var city = "hadera"

	temp := calculate.CalcTemp(city)
	fmt.Printf("The average temperature now at %s is: %d\n", city, temp)
}
