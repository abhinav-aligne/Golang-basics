package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to the user input program"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for the pizza!:")
	
	// comma ok syntax || or error ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ", input)
	fmt.Printf("Type of this rating %T", input)
}
