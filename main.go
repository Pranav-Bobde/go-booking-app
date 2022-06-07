package main

//format package
import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

type UserData struct {
	firstName string
	lastName string
	email string
	userTickets uint
}

var conferenceName = "Go Conference" 
const conferenceTickets int = 50
var remainingTickets uint= 50
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)
var wg = sync.WaitGroup{}
func main() {
	

	greetUsers()

	// for {
		firstName, lastName, email, userTickets := takeUserInputs()
		isValidName, isValidEmail, isValidUserTickets := helper.IsValidInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidUserTickets {
			bookTicket(userTickets, firstName, lastName, email)
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			fmt.Printf("These are all our bookings: %v\n", getFirstNames())

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
			}
		} else{
			if !isValidName{
				fmt.Println("Invalid First or Last Name")
			}
			if !isValidEmail{
				fmt.Println("Invalid Email (does not contain @)")
			}
			if !isValidUserTickets{
				fmt.Println("Invalid tickets entered")
			}
			// continue	
		}
	// }
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v Booking Application! \n", conferenceName)
	fmt.Printf("We have total of %v tickets & %v are still available \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func takeUserInputs() (string, string, string, uint) {
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
	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets
	
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["noOfTickets"] = strconv.FormatUint(uint64(userTickets), 10) 
	// bookings = append(bookings, userData)

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(10*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending Ticket: \n %v \n to email address%v\n", ticket, email)
	wg.Done()
}
