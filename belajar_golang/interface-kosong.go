package main

import "fmt"

func Ups(i int) interface{} {

	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else if i >= 3 {
		return true
	} else {
		return false
	}
}

func main() {
	var data interface{} = Ups(0)
	fmt.Println(data)
}
