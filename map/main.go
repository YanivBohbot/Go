package main

import "fmt"

func main() {

	// var colors map[string]string
	color := map[string]string{
		"red":   "#ff0000",
		"green": "#4df444",
	}
	// color := make(map[string]string)
	// color["wwhite"] = "#fff"

	printMap(color)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("hrx code for", color, "is", hex)
	}
}
