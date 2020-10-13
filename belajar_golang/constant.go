package main

import "fmt"

func main() {
	// Constant merupakan variable yang tidak bisa diubah value nya
	const firstName = "Yusril Fahmi"
	const lastName = "Al Faizi"

	// multiple constant
	const (
		satu = 1
		dua  = 2
	)

	fmt.Println(firstName + lastName)
	fmt.Println(satu + dua)
}
