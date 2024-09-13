package main

import "fmt"

func main() {

	var belanja int
	var kartu, diskon, cashback, buat bool

	fmt.Scan(&belanja, &buat)

	diskon = belanja >= 100000
	kartu = buat && diskon
	cashback = belanja >= 200000 && kartu

	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)

}
