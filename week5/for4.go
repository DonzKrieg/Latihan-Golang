package main

import "fmt"

func main() {

	var n, i, fac int
	fmt.Print("Masukkan bilangan positif: ")
	fmt.Scan(&n)

	fac = 1

	for i = n; i >= 1; i-- {

		fac *= i

	}

	fmt.Println("Hasil pemfaktoran:", fac)

}

// Program Pemfaktoran
// kamus
// 		n, i, fac : integer
// algoritma
//		output("Masukkan Bilangan Positif: ")
// 		input(n)
// 		fac <- 1
// 		for i <- n to 1 do
// 			fac <- fac * i
// 		endfor
// 		output(fac)
// endprogram
