package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// membuat database handle, konfirmasi driver yang digunakan
	db, _ := sql.Open("mysql", "dellis:@/shud")
	defer db.Close()

	// Koneksi dan cek versi server
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to : ", version)
}
