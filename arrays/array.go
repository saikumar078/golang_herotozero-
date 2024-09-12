package main

import "fmt"

func main()  {
	
	var a=[2]int{1,2}
	fmt.Println(a[0],a[1])

	var b=[...]int{1,2,3}
	fmt.Println(b[1])
}