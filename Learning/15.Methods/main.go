package main

import "fmt"

func main() {
	fmt.Println("Methods in Golang")
	abhinav := User{"Abhinav", 22, "abhinav@gmail.com", true}
	fmt.Println(abhinav)
	abhinav.GetStatus()
	abhinav.NewEmail()
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("Is user activate: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@gmail.com"
	fmt.Println("Email of this user is :", u.Email)
}
