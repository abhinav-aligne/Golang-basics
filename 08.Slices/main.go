package main

import "fmt"

func main() {
	fmt.Println("Welcome to the video on slices")

	var fruitList = []string{"Apple", "Banana", "Peaches"}
	fmt.Printf("Type of the fruit list: %T\n", fruitList)
	fmt.Printf("Length of the list: %d\n", len(fruitList))

	// appending an existing list
	fruitList = append(fruitList, "Mango", "Tomatoes")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	
}
