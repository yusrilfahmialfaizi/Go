package main

import "fmt"

func main() {
	var abc [5]int
	abc[0] = 12
	abc[1] = 13
	abc[2] = 14
	abc[3] = 15
	abc[4] = 16

	fmt.Println(abc)

	var cde = [2]int{
		1,
		2,
	}

	fmt.Println(cde)
	fmt.Println(len(cde))

}
