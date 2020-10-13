package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string, age int) {
	firstName = "yusril"
	middleName = "fahmi"
	lastName = "al faizi"
	age = 22

	return firstName, middleName, lastName, age
}

func main() {
	_, _, lastName, Age := getCompleteName()

	fmt.Println(lastName, Age)
}
