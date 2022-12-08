package main

import (
	"fmt"

	"example.com/wunder"
)

func main() {
	fmt.Println("Hello, World from client!")
	var city = "hadera"

	temp := wunder.GetTemp(city)
	fmt.Printf("By wunder temperature now at %s is: %d\n", city, temp)
}
