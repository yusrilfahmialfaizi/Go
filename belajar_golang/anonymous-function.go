package main

import "fmt"

// type declaration func dengan tipedata string dengan return value boolean
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blacklisted : " + name)
	} else {
		fmt.Println("Welcome")
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }
// func blacklistUser(name string) bool {
// 	return name == "user"
// }

func main() {
	blacklist := func(name string) bool {
		return name == "Root"
	}
	registerUser("eko", blacklist)
	// atau bisa langsung dimasukkan menjadi sebuah parameter
	registerUser("Root", func(name string) bool {
		return name == "Root"
	})

}
