package main

import "fmt"

func main() {
	fmt.Println("hello function")
	var nama = printMessage("kurniawan")
	fmt.Println(nama)
}

func printMessage(param string) string {
	fmt.Println("hallo fungsi printMessage")
	fmt.Println("hello params", param)
	return "delnino"
}
