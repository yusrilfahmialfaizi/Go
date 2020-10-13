package main

import "fmt"

type Filter func(string) string

// filter func(string)string adalah parameter dengan sebuah
// function yang parameternya bertipe data string dan return valuenya string
func sayWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

// function dibawah ini sebagai parameter atau argumen dari function di atas
func spamFilter(name string) string {
	if name == "anjing" {
		return "***"
	} else if name == "fucek" {
		return "f**k"
	} else {
		return "baik"
	}
}

func main() {
	spam := spamFilter
	sayWithFilter("anjing", spam)
	sayWithFilter("fucek", spam)
	sayWithFilter("hy", spam)
}
