package main

import "fmt"

// panic function adalah func yang bisa di gunakan untuk menghentikan program
//biasanya dipangggil ketikaterjadi error pada saat program berjalan
// program akan terhenti, teatpi defer masih tetap dieksekusi

// recover adalah function yang bisa digunakan untuk menangkap data panic
// dengan recover proses panic akan berhenti , sehingga program akan tetap berjalan
func endApp() {
	fmt.Println("end")
	message := recover()
	fmt.Println("Terjadi error : ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
