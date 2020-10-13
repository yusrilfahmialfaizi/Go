package main

import "fmt"

func main() {
	var nama string
	var boolean bool

	// cara penulisan seperti diatas
	// atau
	var angka = 0
	// tanpa var boleh menggunakan := pada deklarasi awal saja selnjutnya menggunakan =
	negara := "Indonesia"
	negara = "Amerika"

	nama = "Yusril Fahmi Al Faizi"
	boolean = true

	// multiple variable
	var (
		first  = "First"
		second = "Second"
	)

	fmt.Println(nama)
	fmt.Println(boolean)
	fmt.Println(angka)
	fmt.Println(negara)
	fmt.Println(negara)
	fmt.Println(first)
	fmt.Println(second)

}
