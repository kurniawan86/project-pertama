package main

import (
	"fmt"
)

func main() {
	var avg = calculate(2, 3, 4)
	fmt.Println("average ", avg)

	fmt.Println("hello world")

	var firstname string = "kurniawan"

	var lastname string
	lastname = "nino"
	middle := 90

	fmt.Print("hai\n", firstname, " ", middle, lastname)
	fmt.Println("")

	// coba pointer
	name := new(string)
	// name = "kur"
	fmt.Println(name)
	fmt.Println(*name)
}
