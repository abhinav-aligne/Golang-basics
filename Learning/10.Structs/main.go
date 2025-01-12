package main

import "fmt"

func main() {
	fmt.Println("Welcome to the structs in golang")
	// no inheritenance
	// no super, or parent concpets

	abhinav := User{"Abhinav", "abc@gmail.com", true, 22}
	fmt.Println(abhinav)
	fmt.Printf("Abhinav's details are : %+v\n", abhinav)
	fmt.Printf("Abhinav's age in the struct is  : %+v\n", abhinav.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
