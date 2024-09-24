package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var counter int

	mu := sync.Mutex{}

	increment := func() {
		defer wg.Done()
		mu.Lock()
		counter++
		mu.Unlock()
	}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment()
	}

	decrement := func() {
		defer wg.Done()
		mu.Lock()
		counter--
		mu.Unlock()
	}

	wg.Add(1)
	go decrement()
	wg.Wait()
	fmt.Println(counter)

}
