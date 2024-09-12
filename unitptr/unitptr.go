package main

import (
    "fmt"
    "unsafe"
)

/*
uintptr is an unsigned integer type that is large enough to store the uninterpreted bits of a pointer value.
uintptr is an integer type that is large enough to hold a pointer. In other words, it is an unsigned integer that can store a memory address.
We can use uintptr to get the memory address of a variable.

The full form of uintptr is "unsigned integer pointer".
*/

func main() {
    var x int = 42
    p := &x                     // p is a pointer to x
	fmt.Println(p)
    address := uintptr(unsafe.Sizeof(x))  // Get the address as a number

    fmt.Printf("Address: %v\n", address)   // Print the address as a number
}
