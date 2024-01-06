package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// Ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your lastName: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(
			firstName,
			lastName,
			email,
			userTickets,
			remainingTickets,
		)

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf(
				"Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
				firstName,
				lastName,
				userTickets,
				email,
			)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("The first name or the last name you entered is too short. Please enter atleast two characters.")
			}
			if !isValidEmail {
				fmt.Printf("Email address you entered does not contain @.")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets you entered is invalid.\n")
			}
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf(
		"We have total of %v tickets and %v remaining tickets!\n",
		confTickets,
		remainingTickets,
	)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInputs(
	firstName string,
	lastName string,
	email string,
	userTickets uint,
	remainingTickets uint,
) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
