package main

import "fmt"

const jwtToken = 4000

func main() {

	var username string = "Abhinav"
	fmt.Println("Case-1")
	fmt.Printf("My name is %s\n", username)
	fmt.Printf("Variable is of type : %T\n", username)

	var isLoggedIn bool = false
	fmt.Println("Case-2")
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println("Case-3")
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type : %T \n", smallValue)

	var floatValue float64 = 2557.123456789
	fmt.Println("Case-4")
	fmt.Println(floatValue)
	fmt.Printf("Variable is of type : %T \n", floatValue)

	// Default values and the aliases
	var anotherVariable int
	fmt.Println("Case-5")
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	// Implicit variable declaration
	var website = "www.google.com"
	fmt.Println("Case-6")
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	// Another way of without using 'var' keyword while declaring the variable
	numberOfUsers := 30000
	fmt.Println("Case-7")
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type : %T \n", numberOfUsers)

	// Calling the global constant value inside the 'main' method
	fmt.Println("Case-8")
	fmt.Println(jwtToken)
	fmt.Printf("Variable is of type : %T \n", jwtToken)
}
