// Soal No. 3A

// PseudoCode
// Program Museum1
// kamus
// 		visitor, totVisit, avgVisit, enchance, person, day : real
// 		invalid : boolean
// algoritma
// 		totVisit <- 0
//		enchance <- 0
//		person <- 0
//		day <- 0
// 		repeat
// 			output("Masukkan Jumlah pengunjung: ")
// 			input(visitor)
// 			if visitor >= 0 and visitor <= 200 then
// 				if person = 0 then
// 					enchance <- enchance + 0
// 				else if visitor > person then
// 					enchance <- enchance + 1
// 				else if visitor < person then
// 					enchance <- enchance + o
// 				endif
// 				person <- visitor
//				day <- day + 1
//				totVisit <- totVisit + visitor
// 			endif
// 		until visitor < 0 or visitor > 200
// 		if day = 0 then
// 			day <- 1
// 		endif
// 		avgVisit <- totVisit div day
//		output("Jumlah Peningkatan Pengunjung:", enchance)
//		output("Rata-rata Jumlah Pengunjung:", avgVisit)
// endprogram

// Golang
package main

import "fmt"

func main() {

	var visitor, totVisit, avgVisit, enchance, person, day float32
	var invalid bool

	totVisit = 0
	enchance = 0
	person = 0
	day = 0

	for !invalid {

		fmt.Print("Masukkan Jumlah pengunjung: ")
		fmt.Scan(&visitor)

		if visitor >= 0 && visitor <= 200 {

			if person == 0 {

				enchance = enchance + 0

			} else if visitor > person {

				enchance = enchance + 1

			} else if visitor < person {

				enchance = enchance + 0

			}

			person = visitor
			day = day + 1
			totVisit = totVisit + visitor

		}

		invalid = visitor < 0 || visitor > 200

	}

	if day == 0 {

		day = 1

	}

	avgVisit = totVisit / day
	fmt.Println("Jumlah Peningkatan Pengunjung:", enchance)
	fmt.Println("Rata-rata Jumlah Pengunjung:", avgVisit)

}
