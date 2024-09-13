package main

import "fmt"

type mahasiswa struct {
	nama     string
	nim      int
	eprt     int
	ipk      float64
	semester int
}

const NMAX int = 1000

type tabMhs [NMAX]mahasiswa

func main() {
	var data tabMhs
	var n int

	baca(&data, &n)
	fmt.Println(highestEprt(data, n))
	fmt.Println(lowestIPK(data, n))
	fmt.Println(averageSemester(data, n))
}

func baca(A *tabMhs, n *int) {
	fmt.Scan(n)
	if *n > NMAX {
		*n = NMAX
	}

	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i].nama, &A[i].nim, &A[i].eprt, &A[i].ipk, &A[i].semester)
	}
}

func highestEprt(A tabMhs, n int) int {
	var idx_max int
	var k int

	idx_max = 0
	k = 1

	for k < n {
		if A[idx_max].eprt < A[k].eprt {
			idx_max = k
		}
		k++
	}
	return A[idx_max].eprt
}

func lowestIPK(A tabMhs, n int) float64 {
	var idx_min int
	var k int

	idx_min = 0
	k = 1

	for k < n {
		if A[idx_min].eprt > A[k].eprt {
			idx_min = k
		}
		k++
	}
	return A[idx_min].ipk
}

func averageSemester(A tabMhs, n int) float64 {
	var total int
	var rata float64

	total = 0
	for i := 0; i < n; i++ {
		total += A[i].semester
	}
	rata = float64(total) / float64(n)
	return rata
}
