package main

import "fmt"

func main() {
	/*Declaring a map - Approach 1*/
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#00FF00",
		"white": "#ffffff",
	}

	printMap(colors)

	/*Declaring a map - Approach 2*/
	marvel := make(map[string]string)
	//var marvel map[string]string
	marvel["ironman"] = "Tony Stark"
	marvel["blackwidow"] = "Natasha Romanaf"

	fmt.Println(marvel)

	/*Deleting the contents of a map*/
	delete(marvel, "ironman")
	fmt.Println(marvel)

	/*Iterating a map*/
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " " + hex)
	}
}
