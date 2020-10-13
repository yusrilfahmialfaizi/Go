package main

import "fmt"

func getGoodBye(name string) string {
	return name
}

func main() {
	// menyimpan sebuah function kedalam sebuah variable
	goodbye := getGoodBye
	fmt.Println(goodbye("Good Bye Go language"))
	// Or
	result := goodbye("say goodbye")
	fmt.Println(result)
}
