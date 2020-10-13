package main

import "fmt"

func main() {
	// membuat alias pada tipe data
	type str string

	var nama str = "nama"

	fmt.Println(nama)
}
