package main

import "fmt"

func main() {
	// slice adalah potongan data array
	var bulan = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	slice := bulan[2:4]

	fmt.Println(slice[1])
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice[0] = "Libur"
	bulan2 := append(slice, "Libur Baru")
	fmt.Println(bulan2)
	fmt.Println(bulan)
	// make slice
	newSlice := make([]string, 2, 6)
	newSlice[0] = "0"
	newSlice[1] = "1"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copy slice
	fromSlice := bulan[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

	// array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
}
