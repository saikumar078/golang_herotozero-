package main

import (
	"fmt"
	"time"
)



func hello()  {
	for  i:=0;i<=5;i++{
		time.Sleep(time.Second*2)
		fmt.Println("Hello, World!")
	}
	
}

func bye()  {
	for  i:=0;i<=5;i++{
		time.Sleep(time.Second*1)
		fmt.Println("bye")
	}
	
}


func main()  {
	go hello()
	go bye()

	time.Sleep(time.Second*10)

}