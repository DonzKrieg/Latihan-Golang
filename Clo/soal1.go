package main

import "fmt"

func main() {

	var c, k float64
	fmt.Print("Masukkan derajat suhu dalam satuan Celcius: ")
	fmt.Scan(&c)

	k = c + 273
	fmt.Println("Hasil konversi:", k, "Derajat")

}
