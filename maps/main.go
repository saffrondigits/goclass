package main

import "fmt"

func main() {
	myMap := map[string]string{
		"dave":  "Bangalore",
		"bob":   "London",
		"sai":   "Ohio",
		"james": "Singapore",
	}

	fmt.Println(myMap)

	myMap["david"] = "SF"

	fmt.Println(myMap)

	myMap["dave"] = "Dubai"

	fmt.Println(myMap)

	takValue, ok := myMap["sai"]

	fmt.Println(takValue, ok)

	fmt.Println("Printing the map's key and value")
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}

func PhoneBook() {
	contacts := map[string]string{
		"dave":  "+911542544",
		"bob":   "+446464646",
		"sai":   "+15464",
		"james": "+56161441",
	}

	fmt.Printf("contacts: %v\n", contacts)
}
