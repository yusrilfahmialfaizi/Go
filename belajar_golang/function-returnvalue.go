package main

import "fmt"

func returnValue(name string) string {
	rV := "hello " + name
	return rV
}

// return multiple value
func multiValue() (string, int) {
	return "alfaizi", 22
}

func main() {
	result := returnValue("Al Faizi")
	fmt.Println(result)

	// name, age := multiValue()
	// fmt.Println(name, age)
	// bisa menggunakan _ jika tidak membutuhkan return value yang dipilih
	name, _ := multiValue()
	fmt.Println(name)
}
