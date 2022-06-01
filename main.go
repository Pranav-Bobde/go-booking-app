package main

//format package
import "fmt"

func main() {
	var conferenceName = "conferenceName" 
	const conferenceTickets int = 50
	var remainingTickets uint= 50

	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Printf("We have total of %v tickets & %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name
	println("Enter your first name: ")
	fmt.Scan(&firstName)
	println("Enter your last name: ")
	fmt.Scan(&lastName)
	println("Enter your email: ")
	fmt.Scan(&email)
	println("Enter Number of Tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}