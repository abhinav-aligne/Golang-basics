package main

import (
	"fmt"
	"sort"
)

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

	highScore := make([]int, 4)
	highScore[0] = 233
	highScore[1] = 53
	highScore[2] = 23
	highScore[3] = 29
	// highScore[4] = 31	// error raised
	fmt.Println(highScore)

	highScore = append(highScore, 666, 777, 888)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	fmt.Println("-------------------------------------------------------------")
	// How to remove a value based on the index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
