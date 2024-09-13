// Soal No. 6C

// PseudoCode
// Program My_Tel-U3
// kamus
// 		n, i, poin, gold, plat, silver : integer
// 		valid : boolean
// algoritma
// 		output("Masukkan Jumlah Pengguna: ")
//		input(n)
// 		gold <- 0
//		plat <- 0
//		silver <- 0
// 		valid <- true
// 		for i <- 1 to n do
//			output("Masukkan Poin orang ke- ", i, ": ")
// 			input(poin)
// 			if poin < 50 then
// 				valid <- not valid
// 				repeat
// 					output("=== Masukan Tidak Valid ===")
//					output("Masukkan Poin Anda: ")
//					input(poin)
// 				until poin >= 50
// 			endif
// 			if poin > 200 then
// 				gold <- gold + 1
// 			else if poin >= 100 and poin <= 200 then
// 				plat <- plat + 1
// 			else if poin >= 50 and poin <= 99 then
// 				silver <- silver + 1
// 			endif
// 		endfor
// 		output("Gold User:", gold)
//		output("Platinum User:", plat)
//		output("Silver User:", silver)
// endprogram

// Golang
package main

import "fmt"

func main() {

	var n, i, poin, gold, plat, silver int
	var valid bool
	fmt.Print("Masukkan Jumlah Pengguna: ")
	fmt.Scan(&n)

	gold = 0
	plat = 0
	silver = 0
	valid = true

	for i = 1; i <= n; i++ {

		fmt.Print("Masukkan Poin orang ke- ", i, ": ")
		fmt.Scan(&poin)

		if poin < 50 {

			valid = !valid
			for !valid {

				fmt.Println("=== Masukan Tidak Valid ===")
				fmt.Print("Masukkan Poin orang ke- ", i, ": ")
				fmt.Scan(&poin)
				valid = poin >= 50

			}

		}

		if poin > 200 {

			gold = gold + 1

		} else if poin >= 100 && poin <= 200 {

			plat = plat + 1

		} else if poin >= 50 && poin <= 99 {

			silver = silver + 1

		}

	}

	fmt.Println("Gold User:", gold)
	fmt.Println("Platinum User:", plat)
	fmt.Println("Silver User:", silver)

}
