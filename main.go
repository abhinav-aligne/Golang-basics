package main

import (
	"booking-app/helpers"
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}

var conferenceName = "Go Conference Hall"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

const conferenceTickets = 50

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	var city, firstName, lastName, email string
	var totalTickets uint
	city, firstName, lastName, email, totalTickets = getUserInput()
	var isValidName, isValidEmail = helpers.ValidateUserInput(firstName, lastName, email)

	switch city {
	case "London":
		remainingTickets = 50
		cityTicket(city, remainingTickets, totalTickets)
	case "Singapore":
		remainingTickets = 50
		cityTicket(city, remainingTickets, totalTickets)
	case "New York":
		remainingTickets = 5
		cityTicket(city, remainingTickets, totalTickets)
	case "Berlin":
		remainingTickets = 10
		cityTicket(city, remainingTickets, totalTickets)
	case "India", "Dubai":
		remainingTickets = 15
		cityTicket(city, remainingTickets, totalTickets)
	default:
		fmt.Println("No valid city selected")
		return 
	}

	if isValidName && isValidEmail {

		if totalTickets == 0 {
			fmt.Printf("It's an invalid input. Please try again!\n")
			return

		} else if totalTickets <= remainingTickets {

			bookTicket(totalTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(totalTickets, firstName, lastName, email)

			var firstNames = getFirstNames()
			fmt.Printf("The list of first names in the booking list: %v\n", firstNames)

		} else if totalTickets == remainingTickets {

			fmt.Printf("You are lucky as we only left with %v remaining tickets.\n", remainingTickets)
			fmt.Printf("Our conference is booked and sold out. Come back next year!\n")
			return

		} else if remainingTickets == 0 {
			fmt.Printf("Our conference is booked and sold out. Come back next year!\n")
			return

		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, totalTickets)
			fmt.Printf("Thank you for your time. Please try again!\n")
			return
		}

	} else {
		if !isValidName {
			fmt.Println("First name or last name that you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Your email address doesn't contain @ sign.")
		}
		return
	}
	wg.Wait()
}


func greetUsers() {
	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets from here! for the %v.\n", conferenceName)
}

func getFirstNames() []string {
	var firstNames = []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, string, uint) {
	var city string
	var firstName string
	var lastName string
	var email string
	var totalTickets uint

	fmt.Printf("Enter the city name: `London:50 or Singapore:30 or New York:5 or Berlin:10 or Hong Kong:15 or Dubai:15`: ")
	fmt.Scan(&city)

	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("Enter the number of tickets you wanted to purchase: ")
	fmt.Scan(&totalTickets)

	return city, firstName, lastName, email, totalTickets
}

func cityTicket(city string, remainingTickets uint, total_tickets uint) {
	if total_tickets > remainingTickets {
		fmt.Println("Your input is invalid. Please confirm the remaining tickets available.")
		return
	}
	remainingTickets = remainingTickets - total_tickets
	fmt.Printf("As you selected %v, so the total remaining tickets are: %v\n", city, remainingTickets)
}

func bookTicket(totalTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - totalTickets

	// create a map for a user data
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: totalTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Println()

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, totalTickets, email)
	fmt.Printf("The number of tickets remaining are %v for %v\n", remainingTickets, conferenceName)
}

func sendTicket(total_tickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", total_tickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending Ticket: %v\n Email Address: %v\n", ticket, email)
	fmt.Println("##########")
	wg.Done()
}
