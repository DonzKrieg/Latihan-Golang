// Soal No. 1A

// PseudoCode
// Program Faktor_Persekutuan
// kamus
// 		m, n, i : integer
// algoritma
//		output("Masukkan nilai M: ")
// 		input(m)
//		output("Masukkan nilai N: ")
// 		input(n)
// 		for i <- 1 to m and n do
// 			if m mod i = 0 and n mod i = 0 then
// 				output(i, " ")
// 			endif
// 		endfor
// endprogram

// Golang
package main

import "fmt"

func main() {

	var m, n, i int
	fmt.Print("Masukkan nilai M: ")
	fmt.Scan(&m)
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&n)

	for i = 1; i <= m && i <= n; i++ {

		if m%i == 0 && n%i == 0 {

			fmt.Print(i, " ")

		}

	}

}
