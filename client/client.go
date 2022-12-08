package main

import (
	"fmt"

	"example.com/wunder"
)

func main() {
	fmt.Println("Hello, World from client!")

	message := wunder.Scrape()
	fmt.Println(message)
}
