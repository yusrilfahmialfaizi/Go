package main

import "fmt"

// Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
// harap digunakan dengan bijak saat proses pengembangan

func main() {
	// scope func main
	name := "yusril"
	counter := 0
	// membuat anonymous func
	increment := func() {
		// value name di atas akan direplace oleh closure yang dibwah
		name = "Al Faizi"
		fmt.Println("Increment")
		// counter ini diambil dari scope diatasnya
		counter++

	}
	// tipe data yang berada di scope di bawahnya tidak akan bisa di akses oleh scope di atasnya

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
