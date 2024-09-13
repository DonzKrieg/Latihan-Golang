package main

import "fmt"

func main() {

	var n, i, s, t, k, l float64

	fmt.Print("Masukkan jumlah segitiga: ")
	fmt.Scan(&n)

	for i = 1; i <= n; i++ {

		fmt.Println("Sisi dan tinggi segitiga ke-", i, ": ")
		fmt.Scan(&s, &t)

		k = 3 * s
		l = 0.5 * s * t

		fmt.Println("Nilai Luasnya adalah:", l, "Nilai Kelilingnya adalah:", k)
	}

}
