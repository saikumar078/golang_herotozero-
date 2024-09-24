package main

import (
	"fmt"
	"strconv"
)


func main()  {
	
	var a int=43
	var b float64=float64(a)

	fmt.Println(b)

	fmt.Printf("n%T",b)
	s:=strconv.Itoa(a)
	fmt.Printf("%T",s)
}