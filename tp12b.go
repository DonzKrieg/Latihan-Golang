package main

import "fmt"

const NMAX int = 10

type tabInt [NMAX]int

func main() {

	var data tabInt
	var nData int
	var x2 int
	var idx int

	fmt.Scan(&x2)

	bacaData(&data, &nData)

	cetakData(data, nData)

	fmt.Print("Hasil pencarian: ")
	binarySearch(data, nData, x2, &idx)

	if idx != -1 {
		fmt.Printf("Bilangan %d ditemukan pada posisi ke-%d\n", x2, idx)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	var input int = 0
	for i := 0; i < NMAX; i++ {
		fmt.Scan(&input)

		if i > 0 && input < A[i-1] {
			break
		}
		A[i] = input
		*n++
	}

	if *n > NMAX {
		*n = NMAX
	}
}

func cetakData(A tabInt, n int) {
	fmt.Printf("Data Bilangan: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i])
	}
	fmt.Printf("\n")
}

func binarySearch(A tabInt, n int, x int, idx *int) {
	var midVal int = A[n/2]
	var mid = n / 2
	*idx = -1

	if x == midVal {
		*idx = mid + 1
	}

	if *idx == -1 {
		if x > midVal {
			for i := mid - 1; i < n; i++ {
				if A[i] == x {
					*idx = i + 1
					break
				}
			}
		} else if x < midVal {
			for i := mid + 1; i > 0; i-- {
				if A[i] == x {
					*idx = i + 1
					break
				}
			}
		}
	}

}