// Soal No. 2

// PseudoCode
// Program Faktor_Persekutuan_Terkecil
// kamus
// 		m, n, i, kpk : integer
// algoritma
//		output("Masukkan nilai M: ")
// 		input(m)
//		output("Masukkan nilai N: ")
// 		input(n)
// 		kpk <- m
// 		for kpk mod n not 0 do
// 			kpk <- kpk + m
// 		endfor
// 		output("KPK nya adalah ", kpk)
// endprogram

// Golang
package main

import "fmt"

func main() {

	var m, n, kpk int
	fmt.Print("Masukkan nilai M: ")
	fmt.Scan(&m)
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	kpk = m
	for kpk%n != 0 {

		kpk = kpk + m

	}

	fmt.Print("KPK nya adalah ", kpk)

}
