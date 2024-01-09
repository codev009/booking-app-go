package main

import (
	"fmt"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"

var (
	remainingTickets uint = 50
	bookings              = make([]UserData, 0)
)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(
			firstName,
			lastName,
			email,
			userTickets,
			remainingTickets,
		)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(
				userTickets,
				firstName,
				lastName,
				email,
			)
			sendTicket(
				userTickets,
				firstName,
				lastName,
				email,
			)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("The first name or the last name you entered is too short. Please enter atleast two characters.\n")
			}
			if !isValidEmail {
				fmt.Printf("Email address you entered does not contain @.\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("The number of tickets you entered is invalid.\n")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf(
		"We have total of %v tickets and %v remaining tickets!\n",
		conferenceTickets,
		remainingTickets,
	)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTicket(
	userTickets uint,
	firstName string,
	lastName string,
	email string,
) {
	remainingTickets = remainingTickets - userTickets

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",
		firstName,
		lastName,
		userTickets,
		email,
	)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("############")
}
