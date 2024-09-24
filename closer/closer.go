package main

import (
	"fmt"
)

func s(name string) func() {
	fmt.Println(name + " I am the outer function")
	return func() {
		println(name + " I am the inner function")
	}
}

func main() {
	// Calling the outer and inner function
	s("hello")() // here we are calling the inner function after outer function
	// Output:
	// hello I am the outer function
	// hello I am the inner function

	// Calling only the outer function
	s("bye") // here we are calling the outer function
	// Output:
	// bye I am the outer function

	// Storing the inner function in s1
	s1 := s("s1") // this calls the outer function and returns the inner function
	// Output:
	// s1 I am the outer function

	// Now we call the inner function via s1
	s1() // this will call the inner function
	// Output:
	// s1 I am the inner function
}
