package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			name: "Yusril",
		}
	}
}

func main() {
	// nill = null
	// hanya bisa digunakan di tipe data interface, function, map, slice, pointer dan channel
	nama := NewMap("")
	// fmt.Println(nama)
	if nama == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(nama["name"])
	}
}
