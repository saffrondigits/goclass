package main

import (
	"fmt"
)

// Define a struct for representing a student
type Student struct {
	FullName string         `json:"full_name"`
	Age      int            `json:"age"`
	Courses  []string       `json:"courses"`
	Grades   map[string]int `json:"grades"`
}

func main() {
	// Create a new student instance using the struct
	student1 := Student{
		FullName: "Alice",
		Age:      20,
		Courses: []string{
			"Mathematics",
			"Computer Science",
			"Physics",
		},
		Grades: map[string]int{
			"Mathematics":      95,
			"Computer Science": 88,
			"Physics":          92,
		},
	}

	// Accessing individual fields
	fmt.Println("Name:", student1.FullName)
	fmt.Println("Age:", student1.Age)

	// Accessing slice field
	fmt.Println("\nCourses:")
	for _, course := range student1.Courses {
		fmt.Println("-", course)
	}

	// Accessing map field
	fmt.Println("\nGrades:")
	for course, grade := range student1.Grades {
		fmt.Printf("%s: %d\n", course, grade)
	}
}
