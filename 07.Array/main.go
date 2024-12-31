package main

import "fmt"

func main() {
	fmt.Println("Welcome to the array in golang")

	var fruits [4]string
	fruits[0] = "Apple"
	fruits[1] = "Tomatoes"
	fruits[3] = "Peaches"

	fmt.Println("Give me the fruit list: ", fruits)
	fmt.Println("Give me the fruit list: ", len(fruits))

	var vegList = [3]string{"potatoes", "beans", "mushroom"}
	fmt.Println("Veggie list is ", vegList)
	fmt.Println("The length of the veggie list is ", len(vegList))


}
