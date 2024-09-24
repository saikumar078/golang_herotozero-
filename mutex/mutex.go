// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var n int          // Shared variable
// 	var mtx sync.Mutex // Mutex to control access to 'n'

// 	// Goroutine 1: Increments the value of 'n' by 4
// 	go func() {
// 		mtx.Lock()               // Lock the mutex before modifying 'n'
// 		defer mtx.Unlock()        // Ensure we unlock after we're done
// 		n += 4
// 		fmt.Println("Goroutine 1 - Incremented n by 4:", n)
// 		time.Sleep(1 * time.Second) // Simulate some work
// 	}()

// 	// Goroutine 2: Decrements the value of 'n' by 1
// 	go func() {
// 		mtx.Lock()               // Lock the mutex before modifying 'n'
// 		defer mtx.Unlock()        // Ensure we unlock after we're done
// 		n -= 1
// 		fmt.Println("Goroutine 2 - Decremented n by 1:", n)
// 	}()

// 	// Give the goroutines some time to finish
// 	time.Sleep(2 * time.Second)
// }


package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int
	mu := sync.Mutex{}

	// Increment function
	increment := func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}

	// Decrement function
	decrement := func() {
		defer wg.Done()
		mu.Lock()
		counter--
		mu.Unlock()
	}

	// Increment 4 times
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go increment()
	}

	// Decrement 1 time
	wg.Add(1)
	go decrement()

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final counter value
	fmt.Println("Final counter value:", counter)
}
