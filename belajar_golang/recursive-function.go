package main

import "fmt"

// recursive function adalah function yang memanggil dirinya sendiri
// contoh kasus : Factorial
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// contoh implementasi recursive function dari function di atas
func factorialRecursive(value int) int {
	if value == 1 {
		// harus ada kondisi berhenti seperti ini
		// jika tidak maka looping akan berjalan terus
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	factorial := factorialLoop(10)
	fmt.Println(factorial)

	factorialR := factorialRecursive(10)
	fmt.Println(factorialR)
}
