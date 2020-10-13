package main

import "fmt"

// type assertion merupakan kemampuan merubah tipedata yang ada menjadi tipedata yang diinginkan
// fitur ini seringkali di gunakan ketika bertemu dengan interface kosong

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int) //panic
	// fmt.Println(resultInt)

	// bisa menggunakan switc case agar lebih aman dan tidak panic
	switch value := result.(type) {
	case string:
		fmt.Println("String : ", value)
	case int:
		fmt.Println("int : ", value)
	default:
		fmt.Println("Unknown")

	}
}
