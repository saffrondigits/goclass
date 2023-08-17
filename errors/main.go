package main

import (
	"errors"
	"fmt"
)

func doSomething() error {
	if true {
		return fmt.Errorf("cannot connect to port 1234")
	}

	return nil
}

func dbConnect() error {
	err := doSomething()
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func main() {

}
