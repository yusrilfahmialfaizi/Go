package main

import "fmt"

func main() {
	data := map[string]int{
		"satu": 1,
		"dua":  2,
	}

	fmt.Println(data)
	fmt.Println(data["satu"])
	fmt.Println(data["dua"])
	fmt.Println(len(data))
	delete(data, "dua")
	fmt.Println(data)
}
