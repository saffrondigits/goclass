package main

import (
	"fmt"
)

// Book struct represents a book in the library
type Book struct {
	Title  string
	Author string
	Genres []string
}

func main() {
	// Create a map to store books using ISBN as the key
	library := make(map[string]Book)

	// Add books to the library
	library["978-0061122415"] = Book{
		Title:  "To Kill a Mockingbird",
		Author: "Harper Lee",
		Genres: []string{"Fiction", "Classic", "Coming-of-age"},
	}

	library["978-0141182608"] = Book{
		Title:  "1984",
		Author: "George Orwell",
		Genres: []string{"Fiction", "Dystopian"},
	}

	// Display all books in the library
	fmt.Println("Books in the Library:")
	for isbn, book := range library {
		fmt.Printf("ISBN: %s\n", isbn)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Printf("Genres: %v\n", book.Genres)
		fmt.Println("-----------------------------")
	}

	// Find a book by ISBN
	isbn := "978-0061122415"
	if book, ok := library[isbn]; ok {
		fmt.Printf("Book with ISBN %s:\n", isbn)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
		fmt.Printf("Genres: %v\n", book.Genres)
	} else {
		fmt.Printf("Book with ISBN %s not found\n", isbn)
	}
}
