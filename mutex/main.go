package main

import (
	"fmt"
	"sync"
)

var number int
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go routine1()
		go routine2()
	}

	wg.Wait()
	fmt.Printf("number: %v\n", number)

}

func routine1() {
	defer wg.Done()

	mutex.Lock()
	number++
	mutex.Unlock()
}

func routine2() {
	defer wg.Done()

	mutex.Lock()
	number++
	mutex.Unlock()
}
