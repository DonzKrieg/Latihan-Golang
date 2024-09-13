// package main

// import "fmt"

// const NMAX int = 20

// type tabInt [NMAX]int

// func main() {
// 	var A tabInt
// 	var n int
// 	baca(&A, &n)
// 	cetakElemen(A, n)
// 	cetakInfo(A, n)
// }

// func baca(A *tabInt, n *int) {
// 	*n = 0
// 	for {
// 		var input int
// 		fmt.Scan(&input)
// 		if input <= 0 {
// 			break
// 		}
// 		if *n < NMAX {
// 			A[*n] = input
// 			*n++
// 		}
// 	}
// }

// func cetakElemen(A tabInt, n int) {
// 	fmt.Print("Elemen Array:")
// 	for i := 0; i < n; i++ {
// 		fmt.Printf(" %d", A[i])
// 	}
// 	fmt.Println()
// }

// func nilaiMax(A tabInt, n int) int {
// 	max := A[0]
// 	for i := 1; i < n; i++ {
// 		if A[i] > max {
// 			max = A[i]
// 		}
// 	}
// 	return max
// }

// func nilaiMin(A tabInt, n int) int {
// 	min := A[0]
// 	for i := 1; i < n; i++ {
// 		if A[i] < min {
// 			min = A[i]
// 		}
// 	}
// 	return min
// }

// func cetakInfo(A tabInt, n int) {
// 	max := nilaiMax(A, n)
// 	min := nilaiMin(A, n)
// 	fmt.Println(max)
// 	fmt.Println(min)
// 	fmt.Println(n)
// }

package main

import "fmt"

const NMAX int = 11

type pemain struct {
	nama, noPung, posisi string
	tinggi               int
}

type tabPemain [NMAX]pemain

func main() {
	var A tabPemain
	var n int
	n = 0
	bacaData(&A, &n)
	cetakPemain(A, n)
	max(A, n)
	min(A, n)
}

func bacaData(A *tabPemain, n *int) {
	var p pemain
	for i := 0; i < NMAX; i++ {
		fmt.Scanln(&p.nama, &p.noPung, &p.posisi, &p.tinggi)
		if p.nama == "none" {
			break
		}
		A[i] = p
		*n++
	}
}

func cetakPemain(A tabPemain, n int) {
	fmt.Print("Data Pemain: \n")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %s %d\n", A[i].nama, A[i].noPung, A[i].posisi, A[i].tinggi)
	}
}

func max(A tabPemain, n int) {
	var t pemain = A[0]
	for i := 0; i < n; i++ {
		if t.tinggi < A[i].tinggi {
			t = A[i]
		}
	}
	fmt.Println(t.nama)
}

func min(A tabPemain, n int) {
	var t pemain = A[0]
	for i := 0; i < n; i++ {
		if t.tinggi > A[i].tinggi {
			t = A[i]
		}
	}
	fmt.Println(t.nama)
}
