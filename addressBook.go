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
	var menuOption, id int
	fmt.Printf("\n----CONTACT MENU----\n1.Add Contact\n2.Update Contact\n3.Search By City OR State\n4.Get Count Of Persons of City or State\n5.Delete Contact\n6.Display Contact List\n7.Exit\n")
	fmt.Scanln(&menuOption)
	switch menuOption {
	case 1:
		addContact()
	case 2:
		fmt.Println("Enter ID Of Contact: ")
		fmt.Scanln(&id)
		updateContact(id)
	case 3:
		searchByCityState()
	case 4:
		countByCityState()
	case 5:
		deleteContact()
	case 6:
		fmt.Println(readDataFromDB())
	case 7:
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
	Update or editing the given contact
*/

func updateContact(id int) {
	var newFName, newLName, newAddress, newCity, newState, newPhoneNum, newEmail string
	var editChoice int

	fmt.Println("Select Choice to Edit\n1.First Name\n2.Last Name\n3.Address\n4.City\n5.State\n6.Phone Number\n7.Email-Id")
	fmt.Println("Enter choice: ")
	fmt.Scanln(&editChoice)
	switch editChoice {
	case 1:
		fmt.Println("New First Name: ")
		fmt.Scanln(&newFName)
		result, _ := db.Exec("UPDATE contact SET first_name = ? WHERE id = ?", newFName, id)
		id, _ := result.RowsAffected()
		fmt.Println(id)
	case 2:
		fmt.Println("New Last Name: ")
		fmt.Scanln(&newLName)
		result, _ := db.Exec("UPDATE contact SET last_name = ? WHERE id = ?", newLName, id)
		id, _ := result.RowsAffected()
		fmt.Printf("Successfully updated %v contact..!", id)
	case 3:
		fmt.Println("New Address: ")
		fmt.Scanln(&newAddress)
		result, _ := db.Exec("UPDATE contact SET address = ? WHERE id = ?", newAddress, id)
		id, _ := result.RowsAffected()
		fmt.Printf("Successfully updated %v contact..!", id)
	case 4:
		fmt.Println("New City: ")
		fmt.Scanln(&newCity)
		result, _ := db.Exec("UPDATE contact SET city = ? WHERE id = ?", newCity, id)
		id, _ := result.RowsAffected()
		fmt.Printf("Successfully updated %v contact..!", id)
	case 5:
		fmt.Println("New State: ")
		fmt.Scanln(&newState)
		result, _ := db.Exec("UPDATE contact SET state = ? WHERE id = ?", newState, id)
		id, _ := result.RowsAffected()
		fmt.Printf("Successfully updated %v contact..!", id)
	case 6:
		fmt.Println("New Phone Number: ")
		fmt.Scanln(&newPhoneNum)
		result, _ := db.Exec("UPDATE contact SET phone_number = ? WHERE id = ?", newPhoneNum, id)
		id, _ := result.RowsAffected()
		fmt.Printf("Successfully updated %v contact..!", id)
	case 7:
		fmt.Println("New Email-Id: ")
		fmt.Scanln(&newEmail)
		result, _ := db.Exec("UPDATE contact SET email = ? WHERE id = ?", newEmail, id)
		id, _ := result.RowsAffected()
		fmt.Printf("Successfully updated %v contact..!", id)
	}
}

/*
@ Delete existing contact from DataBase
*/
func deleteContact() {
	var delete_id int
	fmt.Println("Enter ID to Delete Permanently: ")
	fmt.Scanln(&delete_id)

	result, err := db.Exec("DELETE from Contact WHERE id = ?", delete_id)
	if err != nil {
		panic(err.Error())
	}
	id, err := result.RowsAffected()

	fmt.Printf("Successfully Deleted %v Row..!", id)

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
