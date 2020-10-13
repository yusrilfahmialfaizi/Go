package main

import "fmt"

func main() {
	nilai8 := 10000000000
	nilai16 := int16(nilai8)
	nilai32 := int32(nilai8)
	nilai64 := int64(nilai8)
	// apabila length nilai melebihi dari length tipedata maka nilainya akan berubah

	fmt.Println(nilai8)
	fmt.Println(nilai16)
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	var Str = "ABC"
	var e = Str[1]
	var eString = string(e)

	fmt.Println(Str)
	fmt.Println(e)
	fmt.Println(eString)
}
