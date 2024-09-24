package main

import "fmt"


func main()  {
	
defer	func ()  {
		a:=recover()
	if a!=nil{
		fmt.Println("recovered",a)
	}
	}()

	fmt.Println("Hello, World!----1")
	panic("something went wrong!!!")
	

}
