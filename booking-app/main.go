package main

import (
	"fmt"
	"time"
)

const (
	conferenceTickets int = 50
	conferenceName        = "Go Conference"
)

var (
	remainingTickets uint = 50
	bookings         []UserData
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
		firstName, lastName, email, userTickets := getUserInput()
		isValidName := validateName(firstName, lastName)
		isValidEmail := validateEmail(email)
		isValidTicketNumber := validateTicketNumber(userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookings = bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := make([]string, len(bookings))
	for i, booking := range bookings {
		firstNames[i] = booking.firstName
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) []UserData {
	remainingTickets -= userTickets

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings
}

func validateName(firstName string, lastName string) bool {
	return len(firstName) >= 2 && len(lastName) >= 2
}

func validateEmail(email string) bool {
	for _, ch := range email {
		if ch == '@' {
			return true
		}
	}
	return false
}

func validateTicketNumber(userTickets uint) bool {
	return userTickets > 0 && userTickets <= remainingTickets
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("***********************")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("***********************")
}
