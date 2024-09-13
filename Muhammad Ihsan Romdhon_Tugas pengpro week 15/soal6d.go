// Soal No. 6D

// PseudoCode
// Program My_Tel-U4
// kamus
// 		poin, totPoin, gold, plat, silver : integer
// 		valid, user : boolean
// algoritma
// 		gold <- 0
//		plat <- 0
//		silver <- 0
// 		totPoin <- 0
// 		valid <- true
// 		repeat
//			output("Masukkan Poin Pengguna:: ")
// 			input(poin)
// 			if poin < 50 then
// 				valid <- not valid
// 				repeat
// 					output("=== Masukan Tidak Valid ===")
//					output("Masukkan Poin Pengguna: ")
//					input(poin)
// 				until poin >= 50
// 			endif
// 			if poin > 200 then
// 				gold <- gold + 1
// 				totPoin <- totPoin + poin
// 			else if poin >= 100 and poin <= 200 then
// 				plat <- plat + 1
// 				totPoin <- totPoin + poin
// 			else if poin >= 50 and poin <= 99 then
// 				silver <- silver + 1
// 				totPoin <- totPoin + poin
// 			endif
// 		until totPoin >= 500
// 		output("Gold User:", gold)
//		output("Platinum User:", plat)
//		output("Silver User:", silver)
// endprogram

// Golang
package main

import "fmt"

func main() {

	var poin, totPoin, gold, plat, silver int
	var valid, user bool

	gold = 0
	plat = 0
	silver = 0
	totPoin = 0
	valid = true

	for !user {

		fmt.Print("Masukkan Poin Pengguna: ")
		fmt.Scan(&poin)

		if poin < 50 {

			valid = !valid
			for !valid {

				fmt.Println("=== Masukan Tidak Valid ===")
				fmt.Print("Masukkan Poin Pengguna: ")
				fmt.Scan(&poin)
				valid = poin >= 50

			}

		}

		if poin > 200 {

			gold = gold + 1
			totPoin = totPoin + poin

		} else if poin >= 100 && poin <= 200 {

			plat = plat + 1
			totPoin = totPoin + poin

		} else if poin >= 50 && poin <= 99 {

			silver = silver + 1
			totPoin = totPoin + poin

		}

		user = totPoin >= 500

	}

	fmt.Println("Gold User:", gold)
	fmt.Println("Platinum User:", plat)
	fmt.Println("Silver User:", silver)

}
