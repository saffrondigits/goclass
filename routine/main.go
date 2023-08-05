package main

import (
	"fmt"
	"time"
)

func Print1() {
	fmt.Println("Routine 1")
}

func Print2() {
	fmt.Println("Routine 2")
}

func main() {
	go Print1()
	go Print2()

	time.Sleep(time.Second * 1)
}
