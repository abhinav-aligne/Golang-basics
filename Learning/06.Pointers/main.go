package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")
	
	var pointer *int
	fmt.Println("Value of pointer is ", pointer)

	myNumber := 20
	var ptr = &myNumber

	fmt.Println("Value of the new pointer with the reference: ", ptr)
	fmt.Println("Value of the actual pointer: ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value of the variable reference: ", myNumber) 
}
