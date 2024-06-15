package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "FF0000",
		"green": "00FF00",
		"blue":  "0000FF",
	}
	fmt.Println(colors)

	var colors2 map[string]string
	fmt.Println(colors2)

	// this doesn't work because the map is nil
	//colors2["red"] = "FF0000"

	colors3 := make(map[string]string)
	fmt.Println(colors3)
	colors3["red"] = "FF0000"
	fmt.Println(colors3)
	delete(colors3, "red")
	fmt.Println(colors3)

	for color, hex := range colors {
		fmt.Printf("Hex code for %v is %v\n", color, hex)
	}
}
