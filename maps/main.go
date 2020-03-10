package main

import "fmt"

func main() {

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#b4b4b3",
	// }

	colors["black"] = "#000000"
	colors["white"] = "#ffffff"
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex of", color, "is", hex)
	}
}
