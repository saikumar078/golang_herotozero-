// Go program to illustrate how to convert
// bidirectional channel into the
// unidirectional channel
package main

import "fmt"

// Function to send data into the channel (send-only)
func sending(s chan<- string) {
	s <- "GeeksforGeeks"
}

// Function to receive data from the channel (receive-only)
func receiving(r <-chan string) {
	fmt.Println(<-r) // Receive and print the data from the channel
}

func s1(ch chan<-bool)  {
	ch<-true
}

func r2(ch <-chan bool)  {
	fmt.Println(<-ch)
	
}

func main() {

	// Creating a bidirectional channel
	mychanl := make(chan string)
	mychan2 := make(chan bool)

	// Starting the sending Goroutine
	go sending(mychanl)

	// Receiving the value from the channel
	receiving(mychanl)
	go s1(mychan2)
	r2(mychan2)
}
