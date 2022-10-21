package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "R",
		"green": "G",
		"white": "W",
	}
	colors["yellow"] = "Y"
	//
	//var colors map[string]string
	//
	// colors := make(map[string]string)
	// colors["red"] = "R"
	// delete(colors, "red")
	printMap(colors)
	delete(colors, "yellow")
	fmt.Println(colors)
}
func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Short form of", color, "is", hex)
	}
}
