package main

import (
	"fmt"
)

func handlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func divide(a, b int) int {
	defer handlePanic()
	if b == 0 {
		panic("Division by zero")
	}
	return a / b
}

func main() {
	result := divide(10, 0)
	fmt.Println("Result:", result) // This line will not be executed due to the panic and the recovery mechanism
}
