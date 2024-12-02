package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference Hall"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Thank you for booking a ticket for the %v!\n", conferenceName)

	fmt.Println()

	var bookings = []string{}

	for (remainingTickets > 0) && (len(bookings) < 50) {
		var first_name string
		var last_name string
		var email string
		var total_tickets uint
		var city string

		fmt.Printf("Enter the city name: `London or Singapore`: ")
		fmt.Scan(&city)

		var isValidCity = city == "Singapore" || city == "London"

		if isValidCity {
			if city == "Singapore" {
				remainingTickets = 30
				fmt.Printf("As you selected Singapore, so the total remaining tickets are: %v\n", remainingTickets)
			} else if city == "London" {
				remainingTickets = 50
				fmt.Printf("As you selected London, so the total remaining tickets are: %v\n", remainingTickets)
			}
		} else {
			fmt.Printf("You selected invalid input. Please try again.\n")
			continue
		}

		fmt.Printf("Enter your first name: ")
		fmt.Scan(&first_name)

		fmt.Printf("Enter your last name: ")
		fmt.Scan(&last_name)

		fmt.Printf("Enter your email: ")
		fmt.Scan(&email)

		fmt.Printf("Enter the number of tickets you wanted to purchase: ")
		fmt.Scan(&total_tickets)

		var isValidName bool = len(first_name) >= 2 && len(last_name) >= 2
		var isValidEmail bool = strings.Contains(email, "@")

		if isValidName && isValidEmail {

			if total_tickets == 0 {
				fmt.Printf("It's an invalid input. Please try again!\n")
				continue
			} else if total_tickets < remainingTickets {

				remainingTickets = remainingTickets - total_tickets
				bookings = append(bookings, first_name+" "+last_name)

				fmt.Println()

				fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", first_name, last_name, total_tickets, email)
				fmt.Printf("The number of tickets remaining are %v for %v\n", remainingTickets, conferenceName)

				var firstNames = []string{}
				for _, booking := range bookings {
					var names = strings.Fields(booking)
					firstNames = append(firstNames, names[0])
				}
				fmt.Printf("These are all the bookings list: %v\n", bookings)
				fmt.Printf("The list of first names in the booking list: %v\n", firstNames)

				fmt.Println()
			} else if total_tickets == remainingTickets {

				fmt.Printf("You are lucky as we only left with %v remaining tickets.\n", remainingTickets)
				fmt.Printf("Our conference is booked and sold out. Come back next year!\n")
				break

			} else if remainingTickets == 0 {
				fmt.Printf("Our conference is booked and sold out. Come back next year!\n")
				break

			} else {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, total_tickets)
				fmt.Printf("Thank you for your time. Please try again!\n")
				continue
			}

		} else {
			fmt.Printf("Your input data is not invalid. Try again!\n")
			break
		}
	}
}
