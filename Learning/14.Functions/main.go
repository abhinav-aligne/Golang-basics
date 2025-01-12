package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()

	result := adder(3, 4)
	fmt.Printf("Result is %v\n", result)

	proResult := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("Pro Result is %v\n", proResult)
}

func greeter() {
	fmt.Println("Namestey from golang")
}

func adder(x int, y int) int {
	return x + y
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
