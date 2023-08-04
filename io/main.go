package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "example.txt"

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	_, err = file.WriteString("Hello World! Hello World!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	file, err = os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	buffer := make([]byte, 100)

	n, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	fmt.Println("Read from file:", string(buffer[:n]))
}
