package main

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func main() {
	var nilaiAkhir tabInt
	var banyakNilai int
	bacaNilai(&nilaiAkhir, &banyakNilai)
	cetakNilai(nilaiAkhir, banyakNilai)
}

func bacaNilai(NA *tabInt, n *int) {
	fmt.Scanln(n)
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&(*NA)[i])
	}
}

func cetakNilai(NA tabInt, n int) {
	if n == 0 {
		fmt.Println("Array kosong")
		return
	}
	fmt.Print("Nilai yang terdapat pada array:")
	for i := 0; i < n; i++ {
		fmt.Printf(" %d", NA[i])
	}
	fmt.Println()
}
