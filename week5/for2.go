package main

import "fmt"

func main() {

	var n, i, s, k, l int

	fmt.Print("Masukkan jumlah persegi: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {

		fmt.Print("Sisi persegi ke-", i, ": ")
		fmt.Scan(&s)

		k = 4 * s
		l = s * s

		fmt.Println("Nilai Luasnya adalah:", l, "Nilai Kelilingnya adalah:", k)

	}

}
