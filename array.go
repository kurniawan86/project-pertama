package main

import "fmt"

func main() {
	var buah = [4]string{"apel", "anggur", "jeruk", "kiwi"}
	fmt.Println("jumlah element \t", len(buah))
	fmt.Println("isi array ", buah)

	var angka = [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(angka, "len angka = ", len(angka))

	//array multi dimensi
	var array1 = [2][3]int{[3]int{1, 2, 3}, [3]int{3, 2, 1}}
	var array2 = [2][3]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println("len array multidemesi ", len(array2[0]))

	//print dengan for
	for i := 0; i < len(buah); i++ {
		fmt.Println(i, buah[i])
	}

	//print with slice array
	fmt.Println("slice array ", buah[0:2])
}
