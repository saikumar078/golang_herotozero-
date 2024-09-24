package main

import "fmt"

func main() {
    var input string
    fmt.Println("Enter something:")
    fmt.Scanln(&input)
    fmt.Println("You entered:", input)
}
