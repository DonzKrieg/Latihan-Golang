package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {
	var data tabInt
	var nData int
	var x1 int

	fmt.Scan(&x1)

	fmt.Scan(&nData)
	bacaData(&data, &nData)

	cetakData(data, nData)

	fmt.Print("Hasil pencarian: ")
	if sequentialSearch(data, nData, x1) {
		fmt.Printf("Bilangan ditemukan. Terdapat %d bilangan %d.\n", frekuensiBilangan(data, nData, x1), x1)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	fmt.Printf("Data Bilangan: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Printf("\n")
}

func frekuensiBilangan(A tabInt, n, x int) int {
	var f int = 0
	for i := 0; i < n; i++ {
		if A[i] == x {
			f++
		}
	}
	return f
}

func sequentialSearch(A tabInt, n int, x int) bool {
	for i := 0; i < n; i++ {
		if A[i] == x {
			return true
		}
	}
	return false
}
