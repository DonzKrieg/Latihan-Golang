package main

import "fmt"

func main() {

	var name, pass string
	var percobaan int

	percobaan = 0

	fmt.Print("Masukkan name: ")
	fmt.Scan(&name)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&pass)

	for name != "admin" || pass != "admin" {
		fmt.Println("USERNAME ATAU PASSWORD SALAH!!!")
		fmt.Print("Masukkan username: ")
		fmt.Scan(&name)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&pass)
		percobaan++
	}

	fmt.Printf("%d percobaan gagal\n", percobaan)
	fmt.Printf("Login berhasil\n")

}
