package main

import (
	"errors"
	"fmt"
)

func DoSomething() error {
	return errors.New("something went wrong")
}

func main() {
	err := DoSomething()
	if err != nil {
		// If there's an error, print it
		fmt.Println("Error:", err)
	} else {
		// No error
		fmt.Println("Success!")
	}
}
