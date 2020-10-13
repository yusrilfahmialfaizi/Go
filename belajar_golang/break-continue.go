package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		if i >= 10 {
			// akan diberhentikan
			break
		} else {
			fmt.Println(i)
		}
	}

	for a := 1; a <= 20; a++ {
		if a%2 == 0 {
			fmt.Println("Skip : ", a)
			continue
		} else {
			fmt.Println(a)
		}
	}
}
