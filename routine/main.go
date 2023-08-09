package main

import (
	"fmt"
	"time"
)

func sendData(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Sending: ", i)
		ch <- i
		time.Sleep(time.Second)
	}

	close(ch)
}

func receiveData(ch chan int) {
	for data := range ch {
		fmt.Println("Receiving: ", data)
	}
}

func main() {
	ch := make(chan int)

	go receiveData(ch)
	go sendData(ch)

	time.Sleep(time.Second * 6)
}
