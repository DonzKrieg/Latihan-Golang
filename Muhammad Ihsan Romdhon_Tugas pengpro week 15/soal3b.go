// Soal No. 3B

// PseudoCode
// Program Museum2
// kamus
// 		i, visitor, totVisit, day : integer
// 		valid : boolean
// algoritma
// 		day <- 5
//		totVisit <- 0
//		valid <- true
// 		for i <- 1 to day do
// 			output("Masukkan Jumlah pengunjung hari ke-", i, ": ")
// 			input(visitor)
// 			if visitor < 0 or visitor > 200 then
// 				valid <- not valid
// 				repeat
// 					output("=== Masukan Tidak Valid ===")
//					output("Masukkan Jumlah pengunjung hari ke-", i, ": ")
//					input(visitor)
// 				until visitor >= 0 and visitor <= 200
// 			endif
// 			if visitor >= 0 and visitor <= 200 then
// 				totVisit <- totVisit + visitor
// 			endif
// 		endfor
// 		output("Total Jumlah Pengunjung:", totVisit)
// endprogram

// G9olang
package main

import "fmt"

func main() {

	var i, visitor, totVisit, day int
	var valid bool

	day = 5
	totVisit = 0
	valid = true

	for i = 1; i <= day; i++ {

		fmt.Print("Masukkan Jumlah pengunjung hari ke-", i, ": ")
		fmt.Scan(&visitor)

		if visitor < 0 || visitor > 200 {

			valid = !valid

			for !valid {

				fmt.Println("=== Masukan Tidak Valid ===")
				fmt.Print("Masukkan Jumlah pengunjung hari ke-", i, ": ")
				fmt.Scan(&visitor)
				valid = visitor >= 0 && visitor <= 200

			}

		}

		if visitor >= 0 && visitor <= 200 {

			totVisit = totVisit + visitor

		}

	}

	fmt.Println("Total Jumlah Pengunjung:", totVisit)

}
