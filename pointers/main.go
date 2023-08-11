package main

import "fmt"

// * to get value stored at that location
// & for memory location of the variable
func main() {
	var ptr *int

	number := 5

	ptr = &number

	fmt.Println(ptr)
	fmt.Println(&number)

	// Printing the value stored at the memory location
	fmt.Println(*ptr)
	*ptr = 9
	fmt.Println(*ptr)

	str := "New String"
	PrintString(&str)
	fmt.Println("String pointer value", &str)
	fmt.Println(str)
}

func PrintString(str1 *string) {
	*str1 = "Another String"
	fmt.Println("String pointer value", str1)
}
