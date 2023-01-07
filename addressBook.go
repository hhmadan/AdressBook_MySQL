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
	fmt.Printf("\n----CONTACT MENU----\n1.Add Contact\n2.Search By City OR State\n3.Get Count Of Persons of City or State\n4.Display Contact List\n5.Exit\n")
	fmt.Scanln(&menuOption)
	switch menuOption {
	case 1:
		addContact()
	case 2:
		searchByCityState()
	case 3:
		countByCityState()
	case 4:
		fmt.Println(readDataFromDB())
	case 5:
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
@ retrieve data of contacts based on city and state
*/
func searchByCityState() {
	var inputCityStateName string
	var choice int
	var contacts []Contact
	fmt.Printf("1.Search by City\n2.Search by State\n")
	fmt.Scanln(&choice)

	if choice == 1 {
		fmt.Println("Enter Name of City to search People: ")
		fmt.Scanln(&inputCityStateName)

		rows, err := db.Query("SELECT * FROM contact where city = ?", inputCityStateName)
		//rows, err := db.Query("SELECT * FROM contact where city = ' " + inputCityStateName + " ';")
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var person Contact
			if err := rows.Scan(&person.id, &person.FirstName, &person.LastName, &person.Address, &person.City, &person.State, &person.PhoneNumber, &person.Email); err != nil {
				panic(err.Error())
			}
			contacts = append(contacts, person)
		}
		fmt.Println("Available Contacts in city: ", contacts)
	} else {
		fmt.Println("Enter Name of State to search People: ")
		fmt.Scanln(&inputCityStateName)

		rows, err := db.Query("SELECT * FROM contact where state = ?", inputCityStateName)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			var person Contact
			if err := rows.Scan(&person.id, &person.FirstName, &person.LastName, &person.Address, &person.City, &person.State, &person.PhoneNumber, &person.Email); err != nil {
				panic(err.Error())
			}
			contacts = append(contacts, person)
		}
		fmt.Println("Available Contacts in State: ", contacts)
	}
}

/*
@ Getting count for Contacts available in given City or State
*/
func countByCityState() {
	var inputCityStateName string

	fmt.Println("Enter Name of City or State to get People count: ")
	fmt.Scanln(&inputCityStateName)
	rows, err := db.Query("SELECT COUNT(*) FROM contact WHERE city = ? or state = ?", inputCityStateName, inputCityStateName)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var city string
		if err := rows.Scan(&city); err != nil {
			panic(err.Error())
		}
		fmt.Printf("Found Person Count for given City: %s", city)
	}
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
