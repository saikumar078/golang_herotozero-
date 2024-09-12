package main

import "fmt"


type student struct{
	id int
	name string
	age  int
}

func main()  {

	type student1 struct{
		id int
		name string
		age  int
	}
	
	var st1 student

	st1.name="saikumar"
	st1.id=1
	st1.age=20

	fmt.Println(st1)
	var st2 student1

	st2.name="saikumar1"
	st2.id=11
	st2.age=201
	fmt.Println(st2)
}