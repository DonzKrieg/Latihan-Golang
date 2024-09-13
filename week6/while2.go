package main

import "fmt"

func main() {

	var saldo, pengeluaran int

	fmt.Print("Masukkan Saldo: ")
	fmt.Scan(&saldo)

	fmt.Printf("Masukkan pengeluaran: ")
	fmt.Scan(&pengeluaran)

	for pengeluaran != 0 {
		saldo -= pengeluaran

		fmt.Printf("Masukkan pengeluaran: ")
		fmt.Scan(&pengeluaran)
	}
	fmt.Printf("Sisa saldo: %d", saldo)

}
