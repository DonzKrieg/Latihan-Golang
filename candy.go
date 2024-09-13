package main

import (
	"fmt"
)

const nMax int = 1000

type tabInt [nMax]int

func isiArray(arr *tabInt, n int) {
	/* I.S. Data tersedia dalam piranti masukan
	F.S. arr telah terisi sejumlah bilangan */
	var bil int
	for i := 0; i < n; i++ {
		fmt.Scan(&bil)
		arr[i] = bil
	}
}

func pembagianPermen(arr tabInt, n int) tabInt {
	/*mengembalikan array yang berisi jumlah pembagian permen untuk
	setiap murid*/
	var permen tabInt
	var total int

	if n == 1 {
		permen[0] = 1
	} else {
		for i := 0; i < n; i++ { // minimal 1 anak menerima 1 permen
			permen[i] = 1
		}
		total = n
		for i := 1; i < n; i++ {
			if arr[i] > arr[i-1] { //jika i > 0 dan array index ke i lebih besar dari index sebelumnya
				permen[i] = permen[i-1] + 1
			}
			total += permen[i]
		}
		i := n - 2
		for i >= 0 {
			if arr[i] > arr[i+1] {
				permen[i] = max(permen[i], permen[i+1]+1)
			}
			total += permen[i]
			i--
		}
	}
	return permen
}

func permenMinimum(arr tabInt, n int) int {
	/* Mengembalikan permen minimum yang harus dimiliki Raisa agar semua
	murid mendapatkan permen semua dan sesuai dengan ketentuan */
	var total int
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return total
}

func cetak(arr tabInt, n int) {
	/* I.S. array arr berisi sejumlah n bilangan bulat positif
	F.S. array yang berisi pembagian permen untuk setiap murid dan total minimum permen ditampilkan di layar */

	var permen, total int
	permenArr := pembagianPermen(arr, n)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", permenArr[i])
		permen += permenArr[i]
	}
	fmt.Println()
	total = permenMinimum(permenArr, n)
	fmt.Println(total)
}

func main() {
	var n int
	var rank tabInt
	fmt.Scan(&n)
	isiArray(&rank, n)
	cetak(rank, n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
