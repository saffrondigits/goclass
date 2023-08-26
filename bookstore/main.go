package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

var address = "localhost:8080"

func main() {
	r := gin.Default()

	// Add initial dummy data
	addDummyBooks()

	// API to retrieve all books
	r.GET("/books", getAllBooks)

	// API to get a book by ID
	r.GET("/books/:id", getBookByID)

	fmt.Printf("The book store service is running at address: %v\n", address)
	r.Run(address)
}

func addDummyBooks() {
	books = []Book{
		{ID: 1, Title: "Book 1", Author: "Author 1"},
		{ID: 2, Title: "Book 2", Author: "Author 2"},
		{ID: 3, Title: "Book 3", Author: "Author 3"},
		{ID: 4, Title: "Book 4", Author: "Author 4"},
		{ID: 5, Title: "Book 5", Author: "Author 5"},
	}
}

func getAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, book := range books {
		if book.ID == atoi(id) {
			c.JSON(http.StatusOK, book)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func atoi(s string) int {
	n := 0
	for _, c := range []byte(s) {
		n = n*10 + int(c-'0')
	}
	return n
}
