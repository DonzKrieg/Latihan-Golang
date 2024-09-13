// Soal No. 6E

// PseudoCode
// Program My_Tel-U5
// kamus
// 		poin, totPoin, gold, plat, silver, totGold, totPlat, totSilver, avgGold, avgPlat, avgSilver : real
// 		valid, user : boolean
// algoritma
// 		gold <- 0
//		plat <- 0
//		silver <- 0
// 		totPoin <- 0
// 		totGold <- 0
//		totPlat <- 0
//		totSilver <- 0
// 		valid <- true
// 		repeat
//			output("Masukkan Poin Pengguna: ")
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
// 				totGold = totGold + poin
// 				totPoin <- totPoin + poin
// 			else if poin >= 100 and poin <= 200 then
// 				plat <- plat + 1
// 				totPlat = totPlat + poin
// 				totPoin <- totPoin + poin
// 			else if poin >= 50 and poin <= 99 then
// 				silver <- silver + 1
// 				totSilver = totSilver + poin
// 				totPoin <- totPoin + poin
// 			endif
// 		until totPoin >= 500
// 		if gold = 0 then
// 			gold <- 1
// 		endif
// 		if plat = 0 then
// 			plat <- 1
// 		endif
// 		if silver = 0 then
// 			silver <- 1
// 		endif
// 		avgGold <- totGold div gold
//		avgPlat <- totPlat div plat
//		avgSilver <- totSilver div silver
// 		output("Rata-rata Gold User:", avgGold)
//		output("Rata-rata Platinum User:", avgPlat)
//		output("Rata-rata Silver User:", avgSilver)
// endprogram

// Golang
package main

import "fmt"

func main() {

	var poin, totPoin, gold, plat, silver, totGold, totPlat, totSilver, avgGold, avgPlat, avgSilver float64
	var valid, user bool

	gold = 0
	plat = 0
	silver = 0
	totPoin = 0
	totGold = 0
	totPlat = 0
	totSilver = 0
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
			totGold = totGold + poin
			totPoin = totPoin + poin

		} else if poin >= 100 && poin <= 200 {

			plat = plat + 1
			totPlat = totPlat + poin
			totPoin = totPoin + poin

		} else if poin >= 50 && poin <= 99 {

			silver = silver + 1
			totSilver = totSilver + poin
			totPoin = totPoin + poin

		}

		user = totPoin >= 500

	}

	if gold == 0 {

		gold = 1

	}

	if plat == 0 {

		plat = 1

	}

	if silver == 0 {

		silver = 1

	}

	avgGold = totGold / gold
	avgPlat = totPlat / plat
	avgSilver = totSilver / silver

	fmt.Println("Rata-rata Gold User:", avgGold)
	fmt.Println("Rata-rata Platinum User:", avgPlat)
	fmt.Println("Rata-rata Silver User:", avgSilver)

}
