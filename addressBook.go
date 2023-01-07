package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Contact struct {
	id                                                            int
	FirstName, LastName, Address, City, State, PhoneNumber, Email string
}

var db *sql.DB

/*
@ Displaying Menu Options to select Operation to be performed on Adress Book
*/
func menu() {
	var menuOption int
	fmt.Printf("\n----CONTACT MENU----\n1.Add Contact\n2.Display Contact List\n3.Exit")
	fmt.Scanln(&menuOption)
	switch menuOption {
	case 1:
		addContact()
	case 2:
		fmt.Println(readDataFromDB())
	case 3:
		//os.Exit(0)
		return
	default:
		fmt.Println("Invalid Choice")
	}
	menu()
}

/* Feed user inputs to database to add contacts in Address Book */

func addContact() (int64, error) {
	var person Contact

	fmt.Println("Enter First Name: ")
	fmt.Scanln(&person.FirstName)

	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&person.LastName)

	fmt.Println("Enter Address: ")
	fmt.Scanln(&person.Address)

	fmt.Println("Enter City: ")
	fmt.Scanln(&person.City)

	fmt.Println("Enter State: ")
	fmt.Scanln(&person.State)

	fmt.Println("Enter Phone Number: ")
	fmt.Scanln(&person.PhoneNumber)

	fmt.Println("Enter Email-Id: ")
	fmt.Scanln(&person.Email)

	result, err := db.Exec("INSERT INTO contact(first_name, last_name, address, city, state, phone_number, email) VALUES (?, ?, ?, ?, ?, ?, ?)", person.FirstName, person.LastName, person.Address, person.City, person.State, person.PhoneNumber, person.Email)
	if err != nil {
		return 0, fmt.Errorf("add person: %v", err)
	}
	fmt.Println("Contact Added Succesfully..!")
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add person: %v", err)
	}
	return id, nil
}

/*
@ retrieving all the data from database
*/

func readDataFromDB() ([]Contact, error) {
	var contacts []Contact

	rows, err := db.Query("SELECT * FROM contact;")
	if err != nil {
		return nil, fmt.Errorf("error in query all contact: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign record to slice
	for rows.Next() {
		var person Contact
		if err := rows.Scan(&person.id, &person.FirstName, &person.LastName, &person.Address, &person.City, &person.State, &person.PhoneNumber, &person.Email); err != nil {
			return nil, fmt.Errorf("error in query all contact: %v", err)
		}
		contacts = append(contacts, person)
	}
	return contacts, nil
}

func main() {
	fmt.Println("***** WELCOME TO ADDRESS BOOK *****")
	var err error

	db, err = sql.Open("mysql", "root:Hemangi9@root@tcp(127.0.0.1:3306)/AddressBook")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	menu()
}
