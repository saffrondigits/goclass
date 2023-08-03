package main

import (
	"fmt"
)

type Phonebook map[string]string

func main() {
	phonebook := make(Phonebook)

	fmt.Println("Welcome to the Phonebook App!")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. Search Contact")
		fmt.Println("3. Delete Contact")
		fmt.Println("4. Show All Contacts")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice (1-5): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addContact(phonebook)
		case 2:
			searchContact(phonebook)
		case 3:
			deleteContact(phonebook)
		case 4:
			showAllContacts(phonebook)
		case 5:
			fmt.Println("Thank you for using the Phonebook App. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addContact(phonebook Phonebook) {
	var name, number string
	fmt.Print("Enter the contact name: ")
	fmt.Scan(&name)

	// Check if the contact already exists
	_, ok := phonebook[name]
	if ok == true {
		fmt.Println("Contact already exists.")
		return
	}

	fmt.Print("Enter the contact number: ")
	fmt.Scan(&number)

	phonebook[name] = number
	fmt.Printf("Contact '%s' added successfully.\n", name)
}

func searchContact(phonebook Phonebook) {
	var name string
	fmt.Print("Enter the contact name to search: ")
	fmt.Scan(&name)

	number, ok := phonebook[name]
	if ok == false {
		fmt.Printf("Contact '%s' not found.\n", name)
		return
	}

	fmt.Printf("Contact Details: %s - %s\n", name, number)
}

func deleteContact(phonebook Phonebook) {
	var name string
	fmt.Print("Enter the contact name to delete: ")
	fmt.Scan(&name)

	_, ok := phonebook[name]
	if !ok {
		fmt.Printf("Contact '%s' not found.\n", name)
		return
	}

	delete(phonebook, name)
	fmt.Printf("Contact '%s' deleted successfully.\n", name)
}

func showAllContacts(phonebook Phonebook) {
	if len(phonebook) == 0 {
		fmt.Println("Phonebook is empty.")
		return
	}

	fmt.Println("\nAll Contacts:")
	for name, number := range phonebook {
		fmt.Printf("%s - %s\n", name, number)
	}
}
