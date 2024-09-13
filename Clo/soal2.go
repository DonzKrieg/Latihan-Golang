package main

import "fmt"

func main() {

	var pendudukAwal, lahir, imigrasi, mati, emigrasi, total int
	fmt.Print("Masukkan data penduduk awal:")
	fmt.Scan(&pendudukAwal)

	fmt.Print("Masukkan data kelahiran: ")
	fmt.Scan(&lahir)

	fmt.Print("Masukkan data imigrasi: ")
	fmt.Scan(&imigrasi)

	fmt.Print("Masukkan data kematian: ")
	fmt.Scan(&mati)

	fmt.Print("Masukkan data emigrasi: ")
	fmt.Scan(&emigrasi)

	total = pendudukAwal + lahir + imigrasi - mati - emigrasi
	fmt.Println("Total Penduduk:", total)

}
