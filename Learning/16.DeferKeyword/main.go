package main

import "fmt"

func main() {
	defer fmt.Println("Hello world")
	defer fmt.Println("Defer one")
	defer fmt.Println("Defer two")
	myDefer()

	fmt.Println("Defer keyword in golang")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// 4, 3, 2, 1, 0
// defer keyword in golang
// defer two
// defer one
// hello world
