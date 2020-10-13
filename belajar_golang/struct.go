package main

import "fmt"

// struct adlah kumpulan dari field
// struct adalah template atau prototype data
// struct tidak bisa digunakan langsung
type Biodata struct {
	nama, alamat, jenis_kelamin string
	umur                        int
	Dokumen                     string
}

// Struct Method
func (data Biodata) dt() {
	fmt.Println("DATA : ", data)
}

func main() {
	var bio Biodata
	bio.nama = "Nama"
	bio.alamat = "Alamat"
	bio.jenis_kelamin = "Jenis Kelamin"
	bio.umur = 23

	dta := Biodata{Dokumen: "Dokumen"}
	dta.dt()
	// Struct Literals =>
	// atau bisa dengan bio:= Biodata{......,.......,......}
	fmt.Println(bio)
}
