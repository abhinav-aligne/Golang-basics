// 										ASSIGNMENT
// A common use case for Textio is to send birthday messages.
// Complete the main function. It should print: "Happy birthday! You are now 21 years old!".
// Create a string variable messageStart with the text "Happy birthday! You are now"
// Create an integer variable age set to 21
// Create another string variable messageEnd with the text "years old!"
// The provided fmt.Println statement will print the full message on a single line separated by spaces.

package main

import "fmt"

func main() {
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"
	fmt.Println(messageStart, age, messageEnd)
}
