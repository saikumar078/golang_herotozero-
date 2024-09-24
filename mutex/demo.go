package main

import (
	"fmt"
	"sync"
	"time"
)


func main()  {
	var c int
	var mu sync.Mutex

	go func ()  {
		mu.Lock()
		defer mu.Unlock()
		c-=1
		fmt.Println("the goroutine1 decrement by -1",c)

	}()

	go func ()  {
		mu.Lock()
		defer mu.Unlock()
		c+=4
		fmt.Println("the goroutine1 increment by +4",c)
	}()


	
	time.Sleep(2* time.Second)

	fmt.Println("c=",c)
	mu.Lock()

}