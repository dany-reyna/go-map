package main

import "fmt"

func main() {
	var colors map[string]string

	colors2 := make(map[int]string)
	colors2[10] = "#ffffff"

	colors3 := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	fmt.Println(colors)
	fmt.Println(colors2)
	fmt.Println(colors3)

	delete(colors2, 10)
	fmt.Println(colors2)

	printMap(colors3.gitignore)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
