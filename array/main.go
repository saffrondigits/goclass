package main

import "fmt"

func main() {
	var myArr [3]string
	fmt.Printf("myArr: %v\n", myArr)

	fruits := [3]string{"Apple", "Banana", "Orange"}
	fmt.Printf("fruits: %v\n", fruits)

	firstFruit := fruits[0]
	secondFruit := fruits[1]

	fmt.Printf("firstFruit: %v\n", firstFruit)
	fmt.Printf("secondFruit: %v\n", secondFruit)

	for i := 0; i < len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	for index, fruit := range fruits {
		fmt.Printf("Element %d: %s\n", index, fruit)
	}

	fruits[2] = "Mango"
}
