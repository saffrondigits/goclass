package main

import "fmt"

func main() {
	Slice1()
}

func Slice1() {
	var sl []int

	sl = []int{1, 2, 3, 4, 5, 6}

	fmt.Println(sl[0])
	fmt.Println(sl[1])

	sl[2] = 9
	fmt.Printf("sl: %v\n", sl)

	fmt.Printf("sl: %v\n", sl[:])
	fmt.Printf("sl: %v\n", sl[:3])
	fmt.Printf("sl: %v\n", sl[2:])

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

	fmt.Println("using for range loop")
	for i, v := range sl {
		fmt.Println(i, v)
	}
}
