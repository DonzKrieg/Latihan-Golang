package main

import "fmt"

const NMAX = 5

type tabInt struct {
	info [NMAX]int
	n    int
}

func main() {
	var nilaiAkhir tabInt

	for i := 0; i < NMAX; i++ {
		bacaNilai(&nilaiAkhir)
	}

	cetakNilai(nilaiAkhir)
}

func bacaNilai(NA *tabInt) {
	if NA.n < NMAX {
		fmt.Scan(&NA.info[NA.n])
		NA.n++
	}
}

func cetakNilai(NA tabInt) {
	for i := 0; i < NA.n; i++ {
		fmt.Printf("%d ", NA.info[i])
	}
	fmt.Println()
}
