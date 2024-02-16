package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const totalTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0) // its a slice -> variable array

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

// to add wait in threads 
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// infinite loop
	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInputs(firstName, lastName, email, userTickets)
		if isValidName && isValidEmail && isValidTicketNumber {

			bookTickets(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("First names (for privacy purpose) of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("All our tickets are sold for %v.\nPlease come back next year.\n", conferenceName)
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email doesn't have @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Tickets entered are invalid.")
			}
		}
	}
	wg.Wait()

}


func validateUserInputs(firstName string, lastName string, email string, userTickets int) (bool, bool, bool){
	isValidName := (len(firstName) >= 2) && (len(lastName) >= 2)
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := (userTickets > 0) && (userTickets <= int(remainingTickets))
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Println("How many tickets do you need?")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - uint(userTickets)

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: uint(userTickets),
	}

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets.\nWe will send your tickets to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("We have now %v tickets left for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v user", userTickets, firstName, lastName)
	fmt.Println("---------------------------------------")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("---------------------------------------")
	wg.Done() 
}