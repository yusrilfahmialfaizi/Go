package main

import "fmt"

func main() {
	length := 5
	for length <= 10 {
		fmt.Println(length)
		length++

	}

	for i := 0; i <= length; i++ {
		fmt.Print(i)
	}
}
