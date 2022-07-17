package main

import "fmt"

func main() {
	println("hello kondisi")

	var nilai = 4
	if nilai > 10 {
		fmt.Print("oke deh")
	} else if nilai < 5 {
		fmt.Println("not okey")
	} else {
		fmt.Println("noting special")
	}
	// else {
	// 	fmt.Print("not okey")
	// }
}
