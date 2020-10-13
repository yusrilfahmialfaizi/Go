package main

import "fmt"

// defer func adalah func yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function di eksekusi
// defer function akan selalu dieksekusi walaupun terjadi error di function  yang di eksekusi

func login() {
	fmt.Println("Selesai memanggil func login")
}

func runApplication(value int) {
	defer login()
	fmt.Println("menjalankan function setelah function login")
	result := 5 / value
	fmt.Println(result)
}

func main() {
	runApplication(0)
}
