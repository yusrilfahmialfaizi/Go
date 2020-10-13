package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// pointer adalah kemampuan membuat refference ke lokasi data di memory yang sama tanpa menyalin data seperti pass by value akan tetapi pass by refferences
func main() {
	address1 := Address{"Jember", "Jawa Timur", "Indonesia"}
	address2 := address1  //pass by value
	address3 := &address1 //pass by references dengan menambahkan operator &
	var address4 *Address = new(Address)

	address2.City = "Probolinggo"
	address3.City = "Sleman"

	// operator * akan merubah semua value sesuai dengan cariablke yang dipilih
	*address3 = Address{"Jember", "Jawa Timur", "Indonesia"}
	fmt.Println(address1) // address 1nilainya tidak akan berubah karena pass by vlaue
	// artinya value address1 di salin dan dimasukkan ke dalam address2
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)
}
