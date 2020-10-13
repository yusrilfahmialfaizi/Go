package main

import "fmt"

// parameter yang berda di posisi terakhir, memiliki kemampuan dijadikan varargs
// Varargs => datanya bisa menerima lebih dari satu input, atau semacam array
// varargs harus berada dipaling belakand adi parameters
func sumAll(number ...int) int {
	total := 0

	for _, nilai := range number {
		total += nilai
	}

	return total
}
func main() {
	total := sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	total = sumAll(slice1...)
	fmt.Println(total)
}
